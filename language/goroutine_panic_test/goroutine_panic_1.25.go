package main

import (
	"fmt"
	"sync"
)

// 在 goroutine 里开 goroutine，并在第二个 goroutine 里 panic
// 这里使用了 go

var wg sync.WaitGroup

func main() {
	wg.Go(foo1)
	wg.Wait()
	fmt.Println("Main goroutine finished")
}

func foo1() {
	wg.Go(foo2)
}

func foo2() {
	panic("panic in foo2")
}
