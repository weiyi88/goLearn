package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

func loadIni(fileName string, data interface{}) (err error) {
	/*
		参数对校验
		传进来的data参数必须是指针类型（因为需要函数中对其赋值）
		传进来对data参数必须是结构类型指针（因为配置文件中各种键值对需要赋值给结构体的字段）
	*/
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data params should be a pointer") // 新建一个错误
		return
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer") // 新创建一个错误
		return
	}
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// 1,文件得到字节类型数据
	//string(b) // 将字节类型的文件内容转换为字符串
	lineSlice := strings.Split(string(b), "\n")
	// 2,一行一行读数据   index, value
	for idx, line := range lineSlice {
		// 去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		// 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 如果开头是[ 则代表是节
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1) // 新建错误
				return
			}
			// 把这一行首尾[]去掉，取到中间内容，把首尾空格去掉拿内容
			setcionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(strings.TrimSpace(line[1:len(line)-1])) == 0 {
				err = fmt.Errorf("line %d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName 去data 里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if setcionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体，把字段记下来
				}
			}

		}
	}
	// 3,注释就忽略
	// 4,[开头表示节
	// 5,不是[开头就是=分割的键值对
	return
}

func RFile(filePath string) *os.File {
	fileObj, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open fils faild err:%v", err)

	}
	return fileObj
}

func main() {
	var mc MysqlConfig
	err := loadIni("./conf.ini", &mc)
	if err != nil {
		fmt.Printf("load ini faild, errr :%v\n", err)
		return
	}

	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)

	//fileObj := RFile("./conf.ini")
	//var temp = make([]byte, 128)
	//n, err := fileObj.Read(temp)
	//if err != io.EOF {
	//	fmt.Println("文件读完了")
	//}
	//if err != nil {
	//	panic("read file faild ,err:" + err.Error())
	//}
	//fmt.Printf("读取了%d字节数据\n", n)
	//
	//fmt.Println(string(temp[:]))

}
