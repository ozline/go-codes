package code

// 使用 channel 实现交替打印，注意，代码需要在 main 函数里完成
// 否则需要使用 sync.WaitGroup 等待 goroutine 完成

import (
	"fmt"
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
