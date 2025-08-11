package main

import "fmt"

func init() {
	fmt.Println("run init 1 from a.go")
}

func init() {
	fmt.Println("run init 2 from a.go")
}
