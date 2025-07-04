package code

import (
	"bytes"
	"fmt"
	"testing"
)

// nolint
func TestRunSlicePrint(t *testing.T) {
	a := []byte("AAAA/BBBB")
	index := bytes.IndexByte(a, '/')
	b := a[:index]
	c := a[index+1:]
	b = append(b, "CCC"...)
	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))

	// 如果 append 了 3 个C
	// AAAACCCBB
	// AAAACCC
	// CCBB

	// 如果 append 了一堆 C，比如 16 个 C
	// AAAA/BBBB
	// AAAACCCCCCCCCCCCCC
	// BBBB

	// 切片是对底层数组的引用，修改切片会影响原始切片
	// 参考文章：https://blog.csdn.net/weixin_45304503/article/details/138639725
}
