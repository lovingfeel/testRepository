package main

// 下面这两行中各自等号左边和右边的标识符表示同一个代码元素。
// 右边的标识符不是预声明的标识符。
/*
const iota = iota // error: 循环引用
var true = true   // error: 循环引用
*/

var a = b // 可以使用其后声明的变量的标识符
var b = 123

func main() {
	// 下面两行中右边的标识符为预声明的标识符。
	//const iota = iota // ok
	var true = true   // ok
	_ = true

	// 下面几行编译不通过。
	/*
		var c = d // 不能使用其后声明变量标识符
		var d = 123
		_ = c
	*/
}
