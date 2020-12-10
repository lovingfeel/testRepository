package main

import (
	"fmt"
	"unicode/utf8"
)

/**
一种是uint8类型,或者叫byte型,代表了ASCII码的一个字符。
另一种是rune类型,代表一个UTF-8字符。当需要处理中文、日文或者其他复合字符时,
则需要用到rune类型。rune类型实际是一个int32。

len 表示字符串的ASCII 字符个数或字节长度
所以:
ASCII 字符串长度使用len() 长度
Unicode 字符串长度使用utf8.RuneCountInString()
 */
func main() {
	//var a byte = 'a'
	//fmt.Printf("%d %T\n", a, a)
	//var b rune='唐'
	//fmt.Printf("%d %T\n", b, b)
	//fmt.Println(unsafe.Sizeof(b))

	tip := "genji "
	fmt.Println(len(tip))
	tip2 := "褴"
	r, size := utf8.DecodeRuneInString(tip2)
// utf8.MaxRune
	fmt.Println(len(tip2))
	fmt.Printf("%v %v %v",r,size)
}
