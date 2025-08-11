package main

import "fmt"

/*
1. 包下有多个 init()，按照源文件名字典序排序后顺序执行
2. 同一个文件下多个 init，按出现顺序执行
*/

func init() {
	fmt.Println("run init 1 from main.go")
}

func init() {
	fmt.Println("run init 2 from main.go")
}

func main() {
	fmt.Println("run main from main.go")
}
