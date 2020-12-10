package main

import "fmt"

func main() {
	mutex := make(chan struct{}, 1)
	mutex <- struct{}{} // 此行是必需的

	counter := 0
	increase := func() {
		<-mutex // 加锁
		counter++
		mutex <- struct{}{} // 解锁
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

