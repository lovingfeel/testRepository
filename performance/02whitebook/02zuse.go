package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(10)
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		//wg.Wait() // 阻塞在此
	}()
	//wg.Wait() // 阻塞在此
	log.Println(runtime.NumCPU())
}
