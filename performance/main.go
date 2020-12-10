package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	countnum := 640*10000
	//runTask(count)

	intChan := make(chan int, countnum)
	//
	//singePut(countnum,intChan)
	multiPut(countnum,intChan)

}

func runTask(taskCount int)  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("golang并发测试")
	fmt.Printf("processors=%d\n", runtime.NumCPU())
	fmt.Printf("tasks=%d\n", taskCount)
	t1 := time.Now()

	countnum:=  5000*10000
	intChan := make(chan int, countnum)

	for i := 1; i <= countnum; i++ {
		intChan <- i
		//放入数据 insrtyChan(countnum,intChan)
		//go func() {insrtyChan(countnum,intChan) }()//
		//fmt.Println("writeData ", i)
		//time.Sleep(time.Second)
	}

	//for i:=0; i<taskCount; i++ {
	//	//go func() {}()
	//
	//}

	//for runtime.NumGoroutine() > 4 {
	//fmt.Println("current goroutines:", runtime.NumGoroutine())
	//time.Sleep(time.Second)
	//}
	t2 := time.Now()
	fmt.Printf("cost time: %.3fs\n", t2.Sub(t1).Seconds())
	time.Sleep(time.Minute)
	close(intChan) //关闭
}

func singePut(countnum int, intchan chan int)  {
	t1 := time.Now()
	for k := 1; k <500*10000; k++ {
		intchan<- k
	}
	t2 := time.Now()
	fmt.Printf("cost time: %.3fs\n", t2.Sub(t1).Seconds())
}

func multiPut(countnum int, intchan chan int)  {
	start := time.Now()
	exitChan := make(chan bool, 8)

	for rou:=0 ;rou <80;rou++{
		go insertChan(intchan,1,80*10000,exitChan)
	}
	go func(){
		for i := 0; i < 8; i++ {
			<-exitChan
		}

		end := time.Now()
		fmt.Println("使用协程耗时=%.3fs", end.Sub(start).Seconds()  )

		//当我们从exitChan 取出了4个结果，就可以放心的关闭 prprimeChan
	}()
	fmt.Println("main线程退出")
	time.Sleep(time.Minute)

}

func insertChan(intchan chan int,beginnum int,endnum int,exitChan chan bool)  {
	for k := beginnum; k <endnum; k++ {
		intchan<- k
	}
	exitChan<- true
}
