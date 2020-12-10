package main

import (
	"fmt"
	"log"
)

func main()  {
	var i int =0
	b:=i
	log.Println(b)

	for i := 0; i < 3; i++ {
		i := i // 在下面的调用中，左i遮挡了右i。
		// <=> var i = i
		defer func() {
			// 此i为上面的左i，非循环变量i。
			fmt.Println("b:", i)
		}()
	}
}
