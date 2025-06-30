package os

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestGoroutineStack(t *testing.T) {
	// 测试 goroutine 栈分析
	// 这里可以使用 runtime 包来获取当前 goroutine 的栈信息
	buf := make([]byte, 1<<16) // 64KB
	n := runtime.Stack(buf, true)
	fmt.Printf("goroutine stack:\n%s\n", buf[:n])
}

func TestStackGrowth(t *testing.T) {
	// 测试栈增长(主 goroutine)
	// 这里可以通过递归调用来观察栈的增长情况

	stackGrouwth(0, 0)
}

func TestStackGrowthInGoroutine(t *testing.T) {
	// 测试 goroutine 中的栈增长
	// 这里可以通过递归调用来观察 goroutine 的栈增长情况

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		stackGrouwth(0, 0)
		wg.Done() // 完成 goroutine
	}()
	wg.Wait() // 等待 goroutine 完成
}

func stackGrouwth(depth int, lastStack uint64) {
	// var x int
	// fmt.Printf("depth: %d, stack pointer: %p\n", depth, &x)
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	if memStats.StackSys > lastStack {
		fmt.Printf("Triggered stack growth at depth %d: %d KB\n", depth, memStats.StackSys/1024)
	}
	if depth < 1000 {
		stackGrouwth(depth+1, memStats.StackSys)
	}
}

func TestStackGrowthWithInt(t *testing.T) {
	// 测试栈增长
	// 这里可以通过递归调用来观察栈的增长情况

	stackGrouwtWithInt(0, 0)
}

// 携带分配 int 的栈增长测试
func stackGrouwtWithInt(depth int, lastStack uint64) {
	var x int
	fmt.Printf("depth: %d, stack pointer: %p\n", depth, &x)
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	if memStats.StackSys > lastStack {
		fmt.Printf("Triggered stack growth at depth %d: %d KB\n", depth, memStats.StackSys/1024)
	}
	if depth < 1000 {
		stackGrouwtWithInt(depth+1, memStats.StackSys)
	}
}

/*
	某个输出（M1 Pro 32G，MacBook Pro 2021，go1.24）：
	Triggered stack growth at depth 0: 256 KB（可以注意到并非八股里说的从 2-4KB 开始增长）
	Triggered stack growth at depth 2: 288 KB
	Triggered stack growth at depth 5: 320 KB
	Triggered stack growth at depth 10: 384 KB
	Triggered stack growth at depth 22: 512 KB
	Triggered stack growth at depth 44: 768 KB
	Triggered stack growth at depth 88: 1280 KB
	Triggered stack growth at depth 177: 2304 KB
	Triggered stack growth at depth 355: 4352 KB
	Triggered stack growth at depth 712: 8448 KB

	可以注意到第一层就直接扩到 256KB 的栈空间

	对于携带分配的栈增长测试，输出类似，但是我们观察到：
	Triggered stack growth at depth 353: 4352 KB（提前 2 个深度开始分配）
	Triggered stack growth at depth 708: 8448 KB（提前 4 个深度开始分配）

	关于直接分配 256KB，应该是被 runtime 优化了，比如内存空间比较大。

	如果我们将其分在子 goroutine 内，可以观测到：
	Triggered stack growth at depth 0: 320 KB（一开始分的比 256 还大，更违背了八股说的 2-4KB 了）
	Triggered stack growth at depth 2: 352 KB
	Triggered stack growth at depth 5: 384 KB
	Triggered stack growth at depth 10: 448 KB
	Triggered stack growth at depth 22: 576 KB
	Triggered stack growth at depth 44: 832 KB
	Triggered stack growth at depth 88: 1344 KB
	Triggered stack growth at depth 177: 2368 KB
	Triggered stack growth at depth 355: 4416 KB
	Triggered stack growth at depth 712: 8512 KB
*/
