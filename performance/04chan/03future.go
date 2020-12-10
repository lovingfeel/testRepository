package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 3) // 模拟一个工作负载
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	//rand.Seed(time.Now().UnixNano())

	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-a, <-b))
}
