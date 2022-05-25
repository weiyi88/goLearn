package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	str := `{"name":"aring","age":24,"gender":"femal"}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age, p.Gender)
}
