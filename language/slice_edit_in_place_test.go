package code

import (
	"testing"
	"unsafe"
)

/*

利用 go slice 底层原理实现原地修改切片

场景举例：现在有一个切片数组，需要根据里面元素的特性删除某些元素

*/

func TestSliceEditInPlace(t *testing.T) {
	// 例如，我们希望可以去掉 int 数组中的所有奇数元素
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Logf("Original Array: %v", arr)
	// 打印原切片的底层数组地址（用于验证）
	originalAddress := uintptr(unsafe.Pointer(&arr[0]))
	t.Logf("Original Array Address: %x", originalAddress)

	filteredArr := arr[:0] // 创建一个长度为0但容量与arr相同的切片
	for _, v := range arr {
		if v%2 == 1 { // 只保留奇数
			filteredArr = append(filteredArr, v)
		}
	}
	arr = filteredArr
	t.Logf("After filter: %+v", arr)

	// 打印过滤后的切片的底层数组地址
	newAddress := uintptr(unsafe.Pointer(&arr[0]))
	t.Logf("After filter: %+v", arr)
	t.Logf("After filter Array Address: %x", newAddress)

	// 验证底层数组地址是否相同
	if originalAddress == newAddress {
		t.Log("The memory address of the underlying array has not changed.")
	} else {
		t.Error("The memory address of the underlying array has changed!")
	}
}
