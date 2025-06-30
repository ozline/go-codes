package code

// 使用 container/heap 内置库实现最小堆

import (
	"container/heap"
	"fmt"
	"testing"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TestHeap(t *testing.T) {
	h := &MinHeap{2, 1, 5}
	heap.Init(h) // 初始化堆

	heap.Push(h, 3)              // 添加元素
	fmt.Println("Min:", (*h)[0]) // 获取堆顶元素（最小值）

	fmt.Println("Heap elements:")
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h)) // 按顺序弹出元素
	}
}
