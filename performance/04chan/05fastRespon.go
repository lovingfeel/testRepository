package main


import (
	"fmt"
	"math/rand"
	"time"
)

/**
通道里的数据 取完后 会进入assleep状态 ，再强行获取数据会 报错
 */
func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3) + 1
	// 睡眠1秒/2秒/3秒
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()
	c := make(chan int32, 5) // 必须用一个缓冲通道
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	rnd := <- c // 只有第一个回应被使用了
	rnd = <- c
	rnd = <- c

	rnd = <- c
	rnd = <- c
	//rnd = <- c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}
