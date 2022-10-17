package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"monster_name"` // 反射机制
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Sal      float64 `json:"monster_sal"`
	Skill    string  `json:"monster_skill"`
}

// 将struct 序列化
// 成json 字符串
func testStruct() {
	monster := Monster{
		"Aring",
		25,
		"1997-08-08",
		200.22,
		"帅气",
	}
	//fmt.Println(monster)
	data, err := json.Marshal(&monster)
	//fmt.Println(data)
	if err != nil {
		fmt.Printf("序列化错误， err =%v \n", err)
	}
	fmt.Printf("monster 序列化后 = %v\n", string(data))
}

// 反序列化成数据格式，struct，map，切片
func unmarshal() {
	str := "{\"monster_name\":\"Aring\",\"monster_age\":25,\"monster_birthday\":\"1997-08-08\",\"monster_sal\":200.22,\"monster_skill\":\"帅气\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化失败了")
	}
	fmt.Println(monster)
}
func main() {
	//testStruct()
	unmarshal()
}
