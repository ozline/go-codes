package c

import "fmt"

func init() {
	fmt.Println("run init 1 from c.go inside of b package")
}

func init() {
	fmt.Println("run init 2 from c.go inside of b package")
}
