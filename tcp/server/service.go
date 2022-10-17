package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server 端
// 监听端口
// 接收客户端请求建立链接
// 创建goroutine 处理链接
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来端数据", recvStr)
		conn.Write([]byte("收到客户端数据：" + recvStr)) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("listen faild ,err ", err)
		return
	}
	fmt.Println("listen successful ")
	for {
		conn, err := listen.Accept() // 建立链接
		if err != nil {
			fmt.Println("accept faild,err", err)
			continue
		}
		fmt.Println("connect successful ")
		go process(conn)
	}
}
