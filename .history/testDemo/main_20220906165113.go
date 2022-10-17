package main


func addUpper(n int)int {
	res := 0
	for i:=0;i<=n;i++{
		res +=1
	}
	return res
}

func main(){
	// 传统测试方法，在main中使用看看结果
	res := addUpper(10)
	fmt.Println(res)
}