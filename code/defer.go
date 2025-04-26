package code

/*
	学习 defer 的用法，感觉这一个文件就通关 Defer 了

	output:

	0
	0
	3
	3
	0
	4
	0
	5
*/

import (
	"fmt"
)

func RunDeferDemo() {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
	fmt.Println(test4())

	return
}

func test1() (v int) {
	defer fmt.Println(v)
	return v
}

func test2() (v int) {
	defer func() {
		fmt.Println(v)
	}()
	return 3
}

// nolint (标记不检查这个函数的警告)
func test3() (v int) {
	defer fmt.Println(v)
	v = 3
	return 4
}

func test4() (v int) {
	defer func(n int) {
		fmt.Println(n)
	}(v)
	return 5
}
