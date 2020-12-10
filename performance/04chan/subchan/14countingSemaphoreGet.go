package main

import (
	"log"
	"time"
	"math/rand"
)

type Customer struct{id int}
type Bar5 chan Customer

func (bar Bar5) ServeCustomer(c Customer) {
	log.Print("++ 顾客#", c.id, "开始饮酒")
	time.Sleep(time.Second * time.Duration(3 + rand.Intn(16)))
	log.Print("-- 顾客#", c.id, "离开酒吧")
	<- bar // 离开酒吧，腾出位子
}

func main() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar5, 10) // 最对同时服务10位顾客
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second * 2)
		customer := Customer{customerId}
		bar24x7 <- customer // 等待进入酒吧
		go bar24x7.ServeCustomer(customer)
	}
	for {time.Sleep(time.Second)}
}
