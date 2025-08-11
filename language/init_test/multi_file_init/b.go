package main

import "fmt"

func init() {
	fmt.Println("run init 1 from b.go")
}

func init() {
	fmt.Println("run init 2 from b.go")
}
