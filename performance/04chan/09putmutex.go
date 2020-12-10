package main

import "fmt"

func main() {
	mutex := make(chan struct{}, 1) // 容量必须为1

	counter := 0
	increase := func() {
		mutex <- struct{}{} // 加锁
		counter++
		<-mutex // 解锁
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 100000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done; <-done
	fmt.Println(counter) // 2000
}
