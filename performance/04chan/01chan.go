package main

import (
	"fmt"
)

func main() {
	//oneFuc()
	//threeFunc()
	b:=1 << 60
	fmt.Println(b)

}


func oneFuc()  {
		var c chan struct{} // nil
		select {
		case <-c:             // 阻塞操作
		case c <- struct{}{}: // 阻塞操作
		default:
			fmt.Println("Go here.")
		}
}

func twoFunc()  {
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}: // 若此分支被选中，则产生一个恐慌
	case <-c:
	}
}

func threeFunc()  {
	c := make(chan string, 2)
	trySend := func(v string) string{
		select {
		case c <- v: return v
		default: return v // 如果c的缓冲已满，则执行默认分支。
		}
	}
	tryReceive := func() string {
		select {
		case v := <-c: return v
		default: return "-" // 如果c的缓冲为空，则执行默认分支。
		}
	}
	trySend("Hello!") // 发送成功
	trySend("Hi!")    // 发送成功
	fmt.Println(trySend("Bye!"))   // 发送失败，但不会阻塞。
	// 下面这两行将接收成功。
	fmt.Println(tryReceive()) // Hello!
	fmt.Println(tryReceive()) // Hi!
	// 下面这行将接收失败。
	fmt.Println(tryReceive()) // -
}