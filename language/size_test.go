package code

import (
	"fmt"
	"testing"
	"unsafe"
)

/* go 中的字符串实际上是一个结构体

type stringStruct struct {
    data uintptr // 指向底层字节数组的指针
    len  int     // 字符串的长度
}

使用 unsafe.Sizeof 时固定返回 8+8 = 16 字节

*/

func TestStrSize(t *testing.T) {
	var s = "123456789"
	fmt.Printf("len(s)=%d\n  unsafe.Sizeof(s)=%d\n", len(s), unsafe.Sizeof(s))
	s = ""
	fmt.Printf("len(s)=%d\n  unsafe.Sizeof(s)=%d\n", len(s), unsafe.Sizeof(s))
}
