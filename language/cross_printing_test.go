package code

// 使用 channel 实现交替打印，注意，代码需要在 main 函数里完成
// 否则需要使用 sync.WaitGroup 等待 goroutine 完成

import (
	"fmt"
	"sync"
	"testing"
)

func TestCrossPrinting(t *testing.T) {
	numChan := make(chan struct{})
	letterChan := make(chan struct{})

	go func() {
		for i := 1; ; i++ {
			<-numChan
			fmt.Print(i)
			letterChan <- struct{}{}
		}
	}()

	go func() {
		for ch := 'a'; ch <= 'z'; ch++ {
			<-letterChan
			fmt.Print(string(ch))
			numChan <- struct{}{}
		}
		close(numChan)
	}()

	numChan <- struct{}{}

	fmt.Println()
}

// 和上一个函数的区别在于，这个函数逻辑是N 个（N 由用户定义）goroutine 交替打印 1-100 的数字
func TestCrossPrintingN(t *testing.T) {
	num := 100
	n := 50 // 定义 N 个 goroutine，实际 goroutine 取值是 min(num, n)
	wg := sync.WaitGroup{}

	// 创建 N 个 channel
	channels := make([]chan struct{}, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan struct{})
	}

	current := 1
	done := make(chan struct{}) // 用于通知 goroutine 停止

	for i := range n {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for {
				select {
				case <-channels[idx]: // 等待当前 channel 的信号
					if current > num {
						// 如果超过最大值，通知所有 goroutine 停止
						close(done)
						return
					}
					t.Logf("goroutine %d: %d\n", idx+1, current)
					current++
					// 发送信号给下一个 goroutine
					channels[(idx+1)%n] <- struct{}{}
				case <-done: // 如果收到停止信号，退出
					return
				}
			}
		}(i)
	}

	// 启动第一个 goroutine
	channels[0] <- struct{}{}

	// 等待所有 goroutine 完成
	wg.Wait()

	fmt.Println("Finished printing numbers")
}
