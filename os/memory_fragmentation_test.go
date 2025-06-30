package os

import (
	"runtime"
	"testing"
)

func TestMemoryFragmentation(t *testing.T) {
	// 这个测试用例的目的是为了验证内存碎片化的情况
	// 由于 Go 的内存管理和垃圾回收机制，内存碎片化可能会导致内存使用效率降低
	// 这里我们可以通过分配和释放不同大小的内存块来模拟碎片化情况

	var memStats runtime.MemStats

	// 初始化时读取内存统计信息
	runtime.ReadMemStats(&memStats)
	initialHeap := memStats.HeapAlloc
	t.Logf("Initial Heap Allocation: %d KB", initialHeap/1024)

	// 分配多个不同大小的内存块
	for i := 0; i < 100; i++ {
		blockSize := (i + 1) * 1024 // 每次分配递增的内存块
		block := make([]byte, blockSize)
		_ = block // 使用 block 避免编译器优化掉分配

		runtime.ReadMemStats(&memStats)
		t.Logf("Heap Allocation after allocating %d bytes: %d KB", blockSize, memStats.HeapAlloc/1024)
	}

	// 手动触发垃圾回收
	runtime.GC()

	runtime.ReadMemStats(&memStats)
	t.Logf("Heap Allocation after GC: %d KB", memStats.HeapAlloc/1024)
}

// 其实和隔壁的 heap_growth_test.go 很类似，但是这里我们关注的是内存碎片化的情况
// 某个测试中，initialHeap 是 180 KB，分配再 GC 后是 184KB，这个 4KB 的增长可能就是内存碎片化的结果
