package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	// 循环发送20次 Hello World! This is a test demo.
	for i := 0; i < 20; i++ {
		msg := `Hello World! This is a test demo.`
		conn.Write([]byte(msg))
	}
}
