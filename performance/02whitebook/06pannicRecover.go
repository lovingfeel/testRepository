package main

import (
	"fmt"
	"time"
)

/**
1,一个recover函数的返回值为其所恢复的恐慌在产生时被一个panic函数调用所消费的参数。
2,在协程里 defer func recover() 用来抓取恐慌
 */
func main() {
	twoFuc()
}

func oneFuc()  {
	defer func() {
		fmt.Println("正常退出")
	}()
	fmt.Println("嗨！")
	defer func() {
		v := recover()
		fmt.Println("恐慌被恢复了：", v)
	}()
	panic("拜拜！") // 产生一个恐慌
	fmt.Println("执行不到这里")
}

func twoFuc()  {
	fmt.Println("hi!")


	go func() {
		defer func() {
			v := recover()
			fmt.Println("恐慌被恢复了：", v)
		}()

		time.Sleep(time.Second)
		panic(123)
	}()



	for {
		time.Sleep(time.Second)
	}
}