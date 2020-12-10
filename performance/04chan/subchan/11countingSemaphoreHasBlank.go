package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
	"github.com/felixge/fgprof"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomerAtSeat(c int, seat Seat) {
	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
	time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
	bar <- seat // 释放座位，离开酒吧
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	bar24x7 := make(Bar, 10)
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	// 这个for循环和上例不一样。
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		seat := <- bar24x7 // 需要一个空位招待顾客
		go bar24x7.ServeCustomerAtSeat(customerId, seat)
	}
	for {time.Sleep(time.Second)}
}
