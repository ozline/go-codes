package code

import (
	"maps"
	"sync/atomic"
	"time"
)

type CachingMap[K comparable, V any] struct {
	// 仿造 sync.Map 的结构，使用两个 map 轮流存储数据
	index int64
	maps  [2]map[K]V

	// write queue
	queue chan struct {
		key   K
		value V
	}

	getter   func(K) (V, error)
	maxsize  int
	duration time.Duration
}

func NewCachingMap[K comparable, V any](getter func(K) (V, error), maxsize int, duration time.Duration) *CachingMap[K, V] {
	cm := &CachingMap[K, V]{
		index: 0,
		maps: [2]map[K]V{
			make(map[K]V),
			make(map[K]V),
		},
		queue: make(chan struct {
			key   K
			value V
		}, 1024),
		getter:   getter,
		maxsize:  maxsize,
		duration: duration,
	}
	go func(cm *CachingMap[K, V]) {
		duration := cm.duration
		if duration == 0 {
			duration = time.Minute
		}
		ticker := time.NewTicker(duration) // 定时器，用于切换 Map，两次切换中途数据不一致，空窗期大概分钟级
		for {
			select {
			case kv := <-cm.queue:
				cm.maps[(atomic.LoadInt64(&cm.index)+1)%2][kv.key] = kv.value
			case <-ticker.C:
				atomic.StoreInt64(&cm.index, (atomic.LoadInt64(&cm.index)+1)%2)
				if m := cm.maps[(atomic.LoadInt64(&cm.index)+1)%2]; maxsize <= 0 || len(m) <= maxsize {
					maps.Copy(m, cm.maps[atomic.LoadInt64(&cm.index)])
				} else {
					cm.maps[(atomic.LoadInt64(&cm.index)+1)%2] = make(map[K]V)
				}
			}
		}
	}(cm)
	return cm
}

func (cm *CachingMap[K, V]) Get(key K) (value V, ok bool, err error) {
	// fast path, lock-free
	value, ok = cm.maps[atomic.LoadInt64(&cm.index)][key]
	if ok {
		return
	}

	// slow path
	value, err = cm.getter(key)
	if err == nil {
		cm.queue <- struct {
			key   K
			value V
		}{key, value}
	}

	return
}
