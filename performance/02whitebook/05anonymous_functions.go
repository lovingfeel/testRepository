package main
/**
1,闭包内取外部函数的参数的时候是取的地址,而不是调用闭包时刻的参数值.
2,所以我们在使用go func的时候最好把可能改变的值通过值传递的方式传入到闭包之中
   ,避免在协程运行的时候参数值改变导致结果不可预期
 */
import (
	"fmt"
	"time"
)

func main()  {
	//oneFunc()
	twoFunc()
}

func oneFunc()  {
	i := 1

	go func() {
		time.Sleep(100*time.Millisecond)
		fmt.Println("i =", i)
	} ()

	i++
	time.Sleep(1000*time.Millisecond)
}

func twoFunc()  {
	i := 1

	go func(i int) {
		time.Sleep(100*time.Millisecond)
		fmt.Println("i =", i)
	} (i)

	i++
	time.Sleep(1000*time.Millisecond)
}