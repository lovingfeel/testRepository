package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done() // 计数器减1
		wg.Add(1) // 计数值加1
	}()
	wg.Wait() // 主goroutine等待，有可能和第7行并发执行
}
//func main() {
//	var wg sync.WaitGroup
//
//	dosomething(100, &wg) // 调用方法，把计数值加1，并启动任务goroutine
//	dosomething(110, &wg) // 调用方法，把计数值加1，并启动任务goroutine
//	dosomething(120, &wg) // 调用方法，把计数值加1，并启动任务goroutine
//	dosomething(130, &wg) // 调用方法，把计数值加1，并启动任务goroutine
//
//	wg.Wait() // 主goroutine等待，代码逻辑保证了四次Add(1)都已经执行完了
//	fmt.Println("Done")
//}

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	wg.Add(1) // 计数值加1，再启动goroutine
	go func() {
		duration := millisecs * time.Millisecond
		time.Sleep(duration)
		fmt.Println("后台执行, duration:", duration)
		wg.Done()
	}()
}

