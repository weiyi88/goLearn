package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 建立与服务端的链接
// 进行数据收发
// 关闭链接
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 读取用户输入直到换行
		fmt.Println("请输入消息")
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			// 如果输入Q 就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed , err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
