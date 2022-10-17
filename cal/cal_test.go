package cal

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	// 调用
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10)执行错误，期望%v 实际值=%v \n", 55, res)
	}
	// 正确，输出日志
	t.Logf("AddUpper(10)执行正确。。。")
}

//func TestHello(t *testing.T) {
//	fmt.Println("TestHello 被调用了")
//}
