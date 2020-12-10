package main

import (
	"fmt"
	"time"
)
/**
，操作<-time.After(aDuration)将使当前协程进入阻塞状态，
而一个time.Sleep(aDuration)函数调用不会如此。
 */
func AfterDuration(d time.Duration) <- chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

func main() {
	fmt.Println("Hi!")
	//<- AfterDuration(time.Second)
	<-time.After(time.Second)
	fmt.Println("Hello!")
	//<- AfterDuration(time.Second)
	<-time.After(time.Second)
	fmt.Println("Bye!")
}
