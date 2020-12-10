package main

import (
	"fmt"
	"sync"
)

type CounterCopy struct {
	sync.Mutex
	Count int
}

func main() {
	var c CounterCopy
	c.Lock()
	defer c.Unlock()
	c.Count++
	fooCopy(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func fooCopy(c CounterCopy) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
