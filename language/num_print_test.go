package code

import (
	"fmt"
	"sync"
	"testing"
)

func TestNumPrint(t *testing.T) {
	wg := sync.WaitGroup{}
	lock := new(sync.Mutex)
	var a int32 = 0
	var b int32 = 2

	// nolint
	for i := 0; i < 5; i++ {
		go func() {
			if a > b {
				fmt.Println("done")
				return
			}
			lock.Lock()
			defer lock.Unlock()
			a++
			fmt.Printf("i: %d a: %d \n", i, a)
		}()
	}
	wg.Wait()
}

// 问题 1：没有使用 wg.Add(1)，导致这样的 Wait 会永远不会退出
// 问题 2：并发访问 a 和 b 的问题，对 a 的修改并没有正确的被访问，虽然 a++被保护了，但是 if a > b 没有
// 问题 3：i 的闭包问题，在 go 旧版之前，i 是一个引用类型，所有的 goroutine 都会共享这个变量，导致打印的结果是 5，不过在新版已经修改了，但是 i 仍然是混乱的

/*
	i: 0 a: 1
	i: 2 a: 2
	i: 1 a: 3
	i: 3 a: 4
*/
