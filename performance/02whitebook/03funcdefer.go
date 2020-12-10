package main

import (
	"fmt"
	"log"
)

/**
一个协程调用或者延迟调用的实参是在此调用发生时被估值的。更具体地说，
对于一个延迟函数调用，它的实参是在此调用被推入延迟调用堆栈的时候被估值的。
对于一个协程调用，它的实参是在此协程被创建的时候估值的。
一个匿名函数体内的表达式是在此函数被执行的时候才会被逐个估值的，不管此函数是被普通调用还是延迟/协程调用。
 */
func main() {
	//func() {
	//	for i := 0; i < 3; i++ {
	//		defer fmt.Println("a:", i)
	//	}
	//}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			i := i
			log.Println(i)
			defer func() {
				fmt.Println("b:", i)
			}()

			//defer func(i int,j int) {
			//	// 此i为形参i，非实参循环变量i。
			//	fmt.Println("b:", i,j)
			//}(i,i+1)
		}
	}()
}
