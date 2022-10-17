package main

import "fmt"

// 声明接口
type Usb interface {
	// 声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

type Camera struct {
}

type Computer struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

// 接受usb接口的类型变量
// 只要实现了Usb 接口，
func (c Computer) Working(usb Usb) {

	usb.Start()
	usb.Stop()
}

func main() {
	com := Computer{}
	phone := Phone{}
	Camera := Camera{}

	com.Working(phone)
	com.Working(Camera)
}
