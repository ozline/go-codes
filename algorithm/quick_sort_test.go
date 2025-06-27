package code

import "testing"

/*
	快排
*/

// 调用这个函数即可
func TestQuickSort(t *testing.T) {
	arr := []int{10, 7, 8, 9, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	for _, v := range arr {
		t.Log(v)
	}
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high) // 找到第一个划分
		quickSort(arr, low, pi-1)       // 递归左边
		quickSort(arr, pi+1, high)      // 递归右边
	}
}

// partition 的作用是找到划分，这个划分的定义是：左边的元素都小于等于 pivot，右边的元素都大于 pivot
// pivot 我们一般选择最后一个元素
func partition(arr []int, low int, high int) int {
	pivot := arr[high] // 默认pivot 取的是 high 的值
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
