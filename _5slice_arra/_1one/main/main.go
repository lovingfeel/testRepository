package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func main() {
	//num := [...]int{5, 6, 7, 8, 8}
	//
	//fmt.Println("before passing to function ", len(num))
	//changeLocal(num) //num is passed by value
	//fmt.Println("after passing to function ", num)

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}

//func main()  {
//	a := [...]string{"USA", "China", "India", "Germany", "France"}
//	b := a // a copy of a is assigned to b
//	b[0] = "Singapore"
//	fmt.Println("a is ", a)
//	fmt.Println("b is ", b)
//
//
//}

