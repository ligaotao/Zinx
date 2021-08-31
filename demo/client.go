package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("客户端启动")

	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("客户端启动错误")
		return
	}
	for {
		_, err := conn.Write([]byte("hello world"))
		if err != nil {
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("解析错误")
			continue
		}
		fmt.Printf("服务返回 %v", cnt)
		time.Sleep(1 * time.Second)
	}
}
