package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	network := "tcp"
	address := "127.0.0.1:30000"
	listen, err := net.Listen(network, address)
	if err != nil {
		fmt.Printf("main | net.Listen(%s, %s) failed to execute", network, address)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的资源：", recvStr)
	}
}
