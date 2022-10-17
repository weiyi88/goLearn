package main

import (
	"fmt"
	"reflect"
)

// 反射的使用场景
//1，结构体标签使用
//type Demo struct {
//	Name string `json:"name"`
//}
//2，函数适配器

//反射可以在运行时候动态获取变量各种信息，比如变量的类型，类别
//如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
//通过反射，可以修改变量的指，可以调用关联的方法
//使用反射需要，import（"reflect"）

// 变量，interface{}，reflect.Value 是可以相互转换的

// 专门用于反射
// 对类型的操作
func reflectTest01(b interface{}) {
	// 通过反射获取的传入的变量的type，kind，值
	// 1， 现货区到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	//2,获取值
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal = ", rVal)

	// 将rVal 转成interface{}
	iv := rVal.Interface()
	//将 interfave 通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2 = ", num2+100)
}

// 对结构体反射操作
func reflectTest02(b interface{}) {
	//1,获取 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType = ", rTyp)

	//2, 获取reflect.Value
	rVal := reflect.ValueOf(b)

	// 将rVal 转成interface{}
	iv := rVal.Interface()
	fmt.Printf("iv = %v iv=%T \n", iv, iv)

	// 简单类型断言检测
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name = %v \n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

// 类型断言
func TypeJudge(item ...interface{}) {
	for i, x := range item {
		switch x.(type) {
		case bool:
			fmt.Println("bool", i)
		case float64:
			fmt.Println()
		case float32:
			fmt.Println()
		case int, int64:
			fmt.Println()
		case nil:
			fmt.Println()
		case string:
			fmt.Println()
		default:
			fmt.Println()

		}
	}
}

func main() {
	//// 定义一个int
	//var num int = 100
	////var stu struct {
	////	name string
	////	age  int
	////}
	//
	//reflectTest01(num)

	stu := Student{
		"aring",
		25,
	}

	//sta := struct {
	//	Fuck  string
	//	Bitch int
	//}{"fuckfuck", 6666}

	reflectTest02(stu)
}
