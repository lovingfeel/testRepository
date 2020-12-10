package main

import (
	"log"
	"math/rand"
	"time"
)

func ServeCustomer(consumers chan int) {
	for c := range consumers {
		log.Print("++ 顾客#", c, "开始在酒吧饮酒")
		time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
		log.Print("-- 顾客#", c, "离开了酒吧")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const BarSeatCount = 10
	consumers := make(chan int)
	for i := 0; i < BarSeatCount; i++ {
		go ServeCustomer(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}
