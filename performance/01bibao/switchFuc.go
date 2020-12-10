package main

import "fmt"

//默认case 不需要break语句，
func Test() int  {
	return 1
}

func main() {

	//ifelse()
	vartype()
}
//1，case 后面跟一个 变量，switch 写一个带返回值函数
func caseVar()  {
	switch Test() {
	case 1,2:
		fmt.Println("匹配到第一个元素")
	case 3,4:
		fmt.Println("匹配到第两个元素")
	default:
		fmt.Println("什么也没匹配到，走默认值")
	}
	fmt.Println("结束")
}

//switch 后不带表达式
func ifelse()  {
	var n int =20;
	switch  {
	case n==20:
		fmt.Println("等于20")
	case n==30:
		fmt.Println("等于30")
	default:
		fmt.Println("什么也没有，输出默认值")
	}
}

func vartype()  {
	var x interface{}
	var  y string ="10"
	x=y
	// switch 结合 接口.(type) 可以判断实际指向的数据类型 ； 感觉接口变量 可以指向任何数据类型
	switch  x.(type) {
	case int:
		fmt.Println("int 类型")
	case bool,string:
		fmt.Println("布尔和字符串类型")
	default:
		fmt.Println("未知类型")
	}

}