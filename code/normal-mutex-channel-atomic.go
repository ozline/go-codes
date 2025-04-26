package code

// 演示在多协程环境下通过不同方式（普通操作、互斥锁、通道、原子操作）实现并发安全地对共享变量进行减操作，并比较其效果。

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const NUM = 10000

var total int = NUM
var wg sync.WaitGroup
var lock = sync.Mutex{}

func RunNornamlMutexChannelAtomic() {
	Normal()
	total = NUM
	UseMutex()
	total = NUM
	UseChannel()
	UseAtomic(int64(NUM))
}

func Normal() {
	for i := 0; i < NUM; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			total -= 1
		}()
	}
	wg.Wait()
	// 打印一下total的值
	fmt.Println(total)
}

func UseMutex() {
	for i := 0; i < NUM; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			defer wg.Done()
			defer lock.Unlock()
			total -= 1
		}()
	}
	wg.Wait()
	// 打印一下total的值
	fmt.Println(total)
}

func UseChannel() {
	var done = make(chan bool, 1)
	done <- true
	for i := 0; i < NUM; i++ {
		wg.Add(1)
		go func() {
			<-done //多协程下如果取不到会阻塞
			defer wg.Done()
			total -= 1
			done <- true
		}()
	}
	wg.Wait()
	// 打印一下total的值
	fmt.Println(total)
}

func UseAtomic(num int64) {
	temp := atomic.Int64{}
	temp.Store(num)
	for i := 0; i < int(num); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			temp.Add(-1)
		}()
	}
	wg.Wait()
	// 打印一下total的值
	fmt.Println(total)
}
