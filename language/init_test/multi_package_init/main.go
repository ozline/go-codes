package main

import (
	"fmt"

	_ "github.com/ozline/go-codes/language/init_test/multi_package_init/a"
	_ "github.com/ozline/go-codes/language/init_test/multi_package_init/b"
)

func init() {
	fmt.Println("run init 1 from main.go")
}

func init() {
	fmt.Println("run init 2 from main.go")
}

func main() {
	fmt.Println("run main from main.go")
}
