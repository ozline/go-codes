package code

import (
	"fmt"
	"testing"
)

// 即使使用了 []*struct{} 这样的结构，看似是指针传递，但实际上在某种情况下是值传递

func TestSliceByValue(t *testing.T) {
	slice := []*MyStruct{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}}

	printSlice(t, slice)

	// 1. 执行 fooSlice 函数
	fooSlice(slice)
	printSlice(t, slice)

	// 2. 执行 fooSlice3 函数
	slice = fooSlice2(slice)
	printSlice(t, slice)

	// 3. 执行 fooSlice3 函数
	fooSlice3(slice)
	printSlice(t, slice)

	// 4. 执行 fooSlice4 函数
	slice = fooSlice4(slice)
	printSlice(t, slice)
}

func printSlice(t *testing.T, slice []*MyStruct) {
	logStr := ""
	for _, v := range slice {
		logStr += fmt.Sprintf("%d,", v.Value)
	}
	t.Logf("slice: %s", logStr)
}

func fooSlice(slice []*MyStruct) {
	// 删除 slice 中的最后一个元素
	slice = slice[:len(slice)-1]
}

func fooSlice2(slice []*MyStruct) []*MyStruct {
	// 也是删除 slice 中的最后一个元素，但我们会返回一个修改后的 slice
	return slice[:len(slice)-1]
}

func fooSlice3(slice []*MyStruct) {
	// 修改 slice 中的最后一个元素的值
	slice[len(slice)-1].Value = 100
}

func fooSlice4(slice []*MyStruct) []*MyStruct {
	// 修改 slice 中的最后一个元素的值，并返回一个修改后的 slice
	slice[len(slice)-1].Value = 100
	return slice
}
