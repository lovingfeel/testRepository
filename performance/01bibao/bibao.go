package main

import "fmt"

//只会反复调用 匿名函数,n只会初始化一次；
//从man函数得调用可以看出 ，后续是直接调用匿名函数;
// 在java中,n 相当于常量，匿名函数 相当于方法去操作常量
// 使用: 外面的函数可以传递一个值，初始化，后续匿名函数调用可以基于初始化的值和匿名函数传的值做操作

func addUpper() func (int) int {
	var n int =10
	//fmt.Print("heh")
	return func(i int) int {
		n=i+n
		return n
	}
}

func addUpperTwo( n int ) func (int) int {
	//var n int =10
	//fmt.Print("heh")
	return func(i int) int {
		n=i+n
		return n
	}
}

func main() {
	f3:=addUpperTwo(4)
	fmt.Println(f3(1))
	fmt.Println(f3(1))

	f4:=addUpperTwo(7)
	fmt.Println(f4(1))
	fmt.Println(f4(2))

	//f:=addUpper()
	//fmt.Println(f(1))
	//fmt.Println(f(1))
	//
	//
	//f2:=addUpper()
	//fmt.Println("f2",f2(1))
	//fmt.Println("f2",f2(1))
	//
	//fmt.Println(f(1))
}
