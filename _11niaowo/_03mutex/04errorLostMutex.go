package main

import (
	"fmt"
	"sync"
)

func main() {
	//缺少lock或者unlock，两个不是成对出现的
	foo()
}

func foo() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world!")
}
