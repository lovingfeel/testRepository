package main

import (
	"fmt"
	"time"
)

func main() {
	countnum := 6400 * 10000 * 2
	//runTask(count)
	intChan := make(chan int, countnum)
	//单线程写入 耗时3.198s
	//singePut(countnum,intChan)
	//多个携程写入  耗时8.0426872
	multiPut(countnum, intChan)

}

func singePut(countnum int, intchan chan int) {
	t1 := time.Now()
	for k := 1; k < countnum; k++ {
		intchan <- k
	}
	t2 := time.Now()
	fmt.Printf("cost time: %.3fs\n", t2.Sub(t1).Seconds())
}

func multiPut(countnum int, intchan chan int) {
	start := time.Now()
	exitChan := make(chan bool, 16)

	for rou := 0; rou < 16; rou++ {
		go insertChan(intchan, 1, 800*10000, exitChan)
	}
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}

		end := time.Now()
		fmt.Println("使用协程耗时=%.3fs", end.Sub(start).Seconds())
		fmt.Println(len(intchan))
		//当我们从exitChan 取出了4个结果，就可以放心的关闭 prprimeChan
	}()
	fmt.Println("main线程退出")
	time.Sleep(time.Minute)

}

func insertChan(intchan chan int, beginnum int, endnum int, exitChan chan bool) {
	for k := beginnum; k < endnum; k++ {
		intchan <- k
	}
	exitChan <- true
}
