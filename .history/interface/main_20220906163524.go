package main

import "fmt"

type cat struct {
}

type dog struct {
}


type Point struct{
	x int 
	y int 
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func da(x dog) {
	x.speak()
}

func main() {
	//var c cat
	var d dog
	da(d)

}
