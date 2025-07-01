// Code Source：https://github.com/CocaineCong/BiliBili-Code/blob/main/bloomfilter/bloom_filter.go
// Author：CocaineCong
// 仅学习摘抄，代码没有变动

package bloomfilter

import (
	"hash"
	"math"
	"sync"

	"github.com/bits-and-blooms/bitset"
	"github.com/twmb/murmur3"
)

type BloomFilter struct {
	bitset    *bitset.BitSet
	size      uint64 // 位图大小
	hashFuncs []hash.Hash64
	mutex     sync.RWMutex
	// hashMutex sync.Mutex     // 添加专门的哈希函数互斥锁
}

// optimalParameters 计算最优参数，详见 bloomfilter 原理，对于工程师，可以只需要知道是一个数学公式
func optimalParameters(n uint64, p float64) (uint64, int) {
	m := uint64(math.Ceil(-float64(n) * math.Log(p) / (math.Ln2 * math.Ln2)))
	k := int(math.Ceil(math.Ln2 * float64(m) / float64(n)))
	return m, k
}

// NewWithFalsePositiveRate 根据预期元素数量和误判率创建布隆过滤器
func NewWithFalsePositiveRate(expectedItems uint64, falsePositiveRate float64) *BloomFilter {
	// 计算最优位数大小和哈希函数数量
	m, k := optimalParameters(expectedItems, falsePositiveRate)
	funcs := make([]hash.Hash64, k)
	for i := 0; i < k; i++ {
		funcs[i] = murmur3.SeedNew64(uint64(i))
	}
	return &BloomFilter{
		bitset:    bitset.New(uint(m)),
		size:      m,
		hashFuncs: funcs,
		mutex:     sync.RWMutex{},
	}
}

// Add 添加元素到布隆过滤器
func (bf *BloomFilter) Add(item []byte) {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()

	for _, h := range bf.hashFuncs {
		h.Reset()
		_, _ = h.Write(item)
		index := h.Sum64() % bf.size
		bf.bitset.Set(uint(index))
	}
}

// Contains 检查元素是否可能存在布隆过滤器中
func (bf *BloomFilter) Contains(item []byte) bool {
	bf.mutex.RLock()
	defer bf.mutex.RUnlock()

	for _, h := range bf.hashFuncs {
		h.Reset()
		_, _ = h.Write(item)
		index := h.Sum64() % bf.size
		if !bf.bitset.Test(uint(index)) {
			return false // 如果有一个哈希函数对应的位为0，则元素一定不存在
		}
	}
	return true // 所有哈希函数对应的位都为1，可能存在
}
