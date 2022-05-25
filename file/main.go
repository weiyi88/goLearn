package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile() (s string) {
	fileObj, err := os.Open("./test.txt")
	if err != nil {
		panic("open fail err:%v")
	}
	defer fileObj.Close()

	var temp = make([]byte, 128)
	n, err := fileObj.Read(temp)
	if err != io.EOF {
		fmt.Println("文件读完了")
	}
	if err != nil {
		panic("read file faild ,err:" + err.Error())
	}
	fmt.Printf("读取了%d字节数据\n", n)
	return string(temp[:])
}

// 打开文件写内容
func WriteFile(filePaht string) {
	// 判断有无该文件，是，打开只写，追加还是创建
	fileObj, err := os.OpenFile(filePaht, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed , err:%v", err)
		return
	}
	defer fileObj.Close()

	fileObj.Write([]byte("懵逼了\n")) // 写字节
	fileObj.WriteString("这是写字符串")  // 写字符串
}

// 缓存方式写入
func WriteFileBufio(filePath string) {
	fileObj, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed ,err:%v", err)
		return
	}
	defer fileObj.Close()

	// 创建一个写对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello 缓存中")
	wr.Flush() // 缓存写入文件

}

func main() {
	//s := ReadFile()
	//fmt.Println(s)

	//WriteFile("./aaa.txt")
	WriteFileBufio("./bufio.txt")
}
