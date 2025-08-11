package main

import "fmt"

/*
1. 一个文件可以有多个 init()
2. init() 没有参数也没有返回值
3. 不可以主动调用 init() 函数
4. 在初始化全局变量后、main() 函数执行前自动调用

其实可以直接推敲了，这个 init() 函数在编译阶段会被进行做一些优化处理
*/

var initA = printA()

func init() {
	fmt.Println("run init 1")
}

func init() {
	fmt.Println("run init 2")
}

var initB = printB()

func printA() bool {
	fmt.Println("run a")
	return true
}

func printB() bool {
	fmt.Println("run b")
	return true
}

func main() {
	// init() 无法调用
	fmt.Println("run main")
}
