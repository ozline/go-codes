package main

import (
	"fmt"

	_ "github.com/ozline/go-codes/language/init_test/multi_package_init/a"
	_ "github.com/ozline/go-codes/language/init_test/multi_package_init/b"
)

/*
结构：a、b、c 三个包同属于一个库内（同一个 go.mod），其中 c 包在 b 中（但仅仅只是物理位置位于）
导入：main 隐式导入了 a 和 b 包，b 包隐式导入了 c 包

1. 先导入先执行
2. 多个包存在依赖关系，则按依赖关系由里到外执行

可以理解为：main 导入 b 包时，由于 b 包导入了 c 包，所以 c 包的 init 会先被执行
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
