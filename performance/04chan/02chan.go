package main

import (
	"fmt"
)

func main() {
	fibonacci := func() chan uint64 {
		c := make(chan uint64)
		//go func() {
		//	var x, y uint64 = 0, 1
		//	for ; y < (1 << 63); c <- y { // 步尾语句
		//		x, y = y, x+y
		//		fmt.Println(y)
		//
		//		//第一次结束 x=1 ，y=1
		//		//
		//	}
		//	close(c)
		//}()

		go func() {
			var x, y uint64 = 0, 1
			for ; y < (1 << 63);  { // 步尾语句
				x=y
				x, y = y, x+y
				c <- y
				fmt.Println(y)

				//第一次结束 x=1 ，y=1
				//
			}
			close(c)
		}()

		return c
	}

	c := fibonacci()
	a,ok := <-c
	a,ok = <-c
	a,ok = <-c
	a,ok = <-c
	fmt.Println(a,ok)
	//for x, ok := <-c; ok; x, ok = <-c { // 初始化和步尾语句
	//	time.Sleep(time.Second)
	//	fmt.Println("hukkue",x)
	//}
}
