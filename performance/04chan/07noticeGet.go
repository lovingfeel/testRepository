package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{},0)
	// 此信号通道也可以缓冲为1。如果这样，则在下面
	// 这个协程创建之前，我们必须向其中写入一个值。

	go func() {
		fmt.Print("Hello")
		// 模拟一个工作负载。
		time.Sleep(time.Second * 2)

		// 使用一个接收操作来通知主协程。
		//c:= <- done
		fmt.Printf("结果 %T",<- done)
	}()

	done <- struct{}{} // 阻塞在此，等待通知
	fmt.Println(" world!")
}