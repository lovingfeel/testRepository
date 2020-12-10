package main

import "fmt"

func main(){
	m := make(map[int]int)
	m[1] = 1 * 2
	m[2] = 2 * 2
	fmt.Println(m)
	m2 := make(map[string]int)
	m2["python"] = 1
	m2["golang"] = 2
	fmt.Println(m2)
}
