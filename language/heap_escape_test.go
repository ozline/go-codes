package code

import (
	"fmt"
	"testing"
)

// 测试 gc 中的堆逃逸现象
// go test -gcflags="-m" -run=^TestHeapEscape$ ./language
// 会在输出中观察到类似如下的逃逸信息：

/*
language/heap_escape_test.go:15:10: &MyStruct{...} escapes to heap
language/heap_escape_test.go:24:21: t does not escape
language/heap_escape_test.go:26:15: make(map[int]*MyStruct) escapes to heap
language/heap_escape_test.go:29:23: &MyStruct{...} escapes to heap
*/

// 一个函数将结构体存入 map，并返回指针
func processData(id int, cache map[int]*MyStruct) *MyStruct {
	// 创建结构体并取地址
	data := &MyStruct{Value: id}

	// 存入 map（导致逃逸）
	cache[id] = data

	// 返回指针（导致逃逸）
	return data
}

func TestHeapEscape(t *testing.T) {
	// 创建一个 map
	cache := make(map[int]*MyStruct)

	// 调用函数
	result := processData(42, cache)

	// 打印结果
	fmt.Println("Result:", result)
	fmt.Println("Cache:", cache)
}
