package code

import "testing"

// 即使使用了 []*struct{} 这样的结构，看似是指针传递，但实际上在某种情况下是值传递

func TestSliceByValue(t *testing.T) {
	slice := []*MyStruct{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}}

	t.Logf("slice: %+v", slice)

	// 1. 执行 fooSlice 函数
	fooSlice(slice)
	t.Logf("slice: %+v", slice)

	// 2. 执行 fooSlice3 函数
	slice = fooSlice2(slice)
	t.Logf("slice: %+v", slice)
}

func fooSlice(slice []*MyStruct) {
	// 删除 slice 中的最后一个元素
	slice = slice[:len(slice)-1]
}

func fooSlice2(slice []*MyStruct) []*MyStruct {
	// 也是删除 slice 中的最后一个元素，但我们会返回一个修改后的 slice
	return slice[:len(slice)-1]
}
