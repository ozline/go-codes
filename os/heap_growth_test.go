package os

import (
	"fmt"
	"runtime"
	"testing"
)

func TestHeapGrowth(t *testing.T) {
	var memStats runtime.MemStats

	// 初始化时读取内存统计信息
	runtime.ReadMemStats(&memStats)
	initialHeap := memStats.HeapAlloc
	fmt.Printf("Initial Heap Allocation: %d KB\n", initialHeap/1024)

	// 分配大量内存，观察堆的增长
	var data [][]byte
	for i := 0; i < 100; i++ {
		// 每次分配 1MB 的内存
		block := make([]byte, 1<<20) // 1MB
		data = append(data, block)

		// 读取内存统计信息
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap Allocation after %d MB: %d KB\n", (i + 1), memStats.HeapAlloc/1024)
	}

	// 清空分配的数据
	data = nil

	// 手动触发垃圾回收
	runtime.GC()

	// 再次读取内存统计信息
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Heap Allocation after GC: %d KB\n", memStats.HeapAlloc/1024)
}

/*
	Initial Heap Allocation: 179 KB
	Heap Allocation after 1 MB: 1203 KB(可以注意到1203-179 刚好等于 1024KB，也就是 1MB，这个符合预期)
	Heap Allocation after 2 MB: 2227 KB
	Heap Allocation after 3 MB: 3251 KB
	Heap Allocation after 4 MB: 4280 KB
	Heap Allocation after 5 MB: 5307 KB
	Heap Allocation after 6 MB: 6330 KB
	Heap Allocation after 7 MB: 7353 KB
	Heap Allocation after 8 MB: 8373 KB
	Heap Allocation after 9 MB: 9398 KB
	Heap Allocation after 10 MB: 10421 KB
	Heap Allocation after 11 MB: 11445 KB
	Heap Allocation after 12 MB: 12469 KB
	/.../
	Heap Allocation after 98 MB: 100542 KB
	Heap Allocation after 99 MB: 101566 KB
	Heap Allocation after 100 MB: 102590 KB（也是 1024 的增长）
	Heap Allocation after GC: 189 KB（GC 后变 189KB，仍然比原始多 10KB）

	但是这不能说明 10KB 就是内存碎片，也有可能是 GC 为了后续快速分配而保留的内存。
*/
