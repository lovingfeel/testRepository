package main

import "fmt"

func main() {
	//var ids []int
	var ids =make([]int,2,10)
	appendSlice(ids)
	fmt.Printf("%p\n",ids)
	fmt.Println("main", len(ids))
}

func appendSlice(ids []int)   {
	for i := 0; i < 4; i++ {
		ids = append(ids, i)
	}
	fmt.Println("appendSlice", len(ids))
	fmt.Printf("%p\n",ids)
	//return ids
}


