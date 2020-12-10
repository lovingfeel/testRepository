package main

import (
	"log"
	"math/rand"
	"time"
)

/**
1,当一个程序的主协程退出后，此程序也就退出了，即使还有一些其它协程在运行。
2,因是log标准库中的打印函数是经过了同步处理的（下一节将解释什么是并发同步）
3,一个处于睡眠中的（通过调用time.Sleep）或者在等待系统调用返回的协程被认为是处于运行状态，而不是阻塞状态。
4,睡眠和等待系统调用返回子状态被认为是运行状态，而不是阻塞状态。
5，一个是正常的函数调用堆栈。在此堆栈中，相邻的两个调用存在着调用关系。晚进入堆栈的调用被早进入堆栈的调用所调用。 此堆栈中最早被推入的调用是对应协程的启动调用。
另一个堆栈是上面提到的延迟调用堆栈。处于延迟调用堆栈中的任意两个调用之间不存在调用关系。

执行方式“正常的函数调用堆栈”代码是从上往下执行，“延迟调用堆栈” 代码相反从小往上

*/
func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // 睡眠片刻（随机0到2.5秒）
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	time.Sleep(2 * time.Second)
}
