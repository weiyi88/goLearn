package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 3000,
	})
	if err != nil {
		fmt.Println("listen faild :", err)
	}
	defer listen.Close()

	// 阻塞等待数据
	for true {
		var data [1024]byte

		// 接收数据
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed,err:", err)
			continue
		}
		fmt.Printf("data:%v add:%v count:%v", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(append(data[:n], []byte("这是客户端数据")...), addr) //发送数据
		if err != nil {
			fmt.Println("write to udp failed ,err:", err)
			continue
		}
	}

}
