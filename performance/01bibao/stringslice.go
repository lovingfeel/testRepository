package main

import "fmt"

func main() {
	//sliceone()
	changeStr()
}

func sliceone()  {
	str:="123456789"
	//string本质是个切片, 底层有一个字节数组,所以可以用切片截取字符串以及其他的切片操作
	strslice:= str[0:1]
	fmt.Println(strslice)
}

//修改字符串
func changeStr()  {
	str:="354911379@qq.com"
//byte()是按字节来处理，所以修改不了汉字，rune按字符处理
//	arr1:=[]byte(str)
	arr1:=[]rune(str)
	arr1[1]='唐'
	arr1[2]='良'
	str=string(arr1)
	fmt.Println(str)
}