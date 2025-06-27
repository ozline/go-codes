package code

// 优雅退出

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 优雅关闭-多个发送方
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				ch <- id*10 + j
			}
		}(i)
	}

	// 用一个通用 goroutine 关闭 channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 接收数据
	for v := range ch {
		fmt.Println(v)
	}
}

// 优雅关闭-仅一个发送方
func TestGracefulShutdown(t *testing.T) {
	ch := make(chan int)

	// 启动一个 goroutine 发送数据
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	// 接收数据
	for v := range ch {
		println(v)
	}

	fmt.Println("Done")
}
