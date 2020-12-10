package main

import "fmt"

func main()  {
	var a [16]byte
	var s []int
	var m map[string]int

	fmt.Println(a == a)   // true
	fmt.Println(m == nil) // true
	fmt.Println(s == nil) // true
	fmt.Println(nil == map[string]int{}) // false
	fmt.Println(nil == []int{})          // false

	// 下面这些行编译不通过。
	/*
		_ = m == m
		_ = s == s
		_ = m == map[string]int(nil)
		_ = s == []int(nil)
		var x [16][]int
		_ = x == x
		var y [16]map[int]bool
		_ = y == y
	*/
}
