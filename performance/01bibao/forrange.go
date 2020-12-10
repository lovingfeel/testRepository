package main

import "fmt"

type mystruct struct {
	name int
	id   int
}

func main() {
	//在一个数组或切片组合字面量中，如果一个元素的索引下标缺失，
	// 则编译器认为它的索引下标为出现在它之前的元素的索引下标加一。

	//1.字符串切片
	//hearo:=[]string{"tanglw","liubei"}

	//hearo:=[]string{0:"tanglw",1:"liubei"}
	//hearo:=[]string{1:"liubei",0:"tanglw"}
	//hearo:=[]string{0:"liubei","tanglw"}
	//2，结构体切片
	//hearo:=[]mystruct{mystruct{6,7},
	//	{6,7},}

	hearo:=map[mystruct]map[string]int{{1,2}:{"tang":7},
		{3,4}:{"liang":7},}

	for index,value:= range hearo {
		fmt.Println(index,value)
	}
}
