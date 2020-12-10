package main

import (
	"fmt"
	"time"
)

func main() {
	process(time.Second)
}

func process(timeout time.Duration) bool {
	ch := make(chan bool, 1)

	go func() {
		// 模拟处理耗时的业务
		time.Sleep((timeout + time.Second))
		ch <- true // block
		fmt.Println("exit goroutine")
	}()
	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return false
	}
}
