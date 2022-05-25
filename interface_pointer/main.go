package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(foot string) {
	fmt.Println("猫吃%s。。。\n", foot)
}

func main() {
	var a1 animal
	//c1 := cat{"tom", 4}
	c2 := &cat{"loser", 4}

	//a1 = c1
	a1 = c2

	fmt.Println(a1)
}
