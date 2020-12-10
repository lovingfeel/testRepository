package main
import (
	"net/http"
	_ "net/http/pprof"
)
import (
	"log"
	"math/rand"
	"time"
)
type Seat2 int
type Bar2 chan Seat2

func (bar Bar2) TenServeCustomerAtSeat(consumers chan int) {
	for c := range consumers {
		seatId := <- bar
		log.Print("++ 顾客#", c, "在第", seatId, "个座位开始饮酒")
		time.Sleep(time.Second * time.Duration(2 + rand.Intn(6)))
		log.Print("-- 顾客#", c, "离开了第", seatId, "个座位")
		bar <- seatId // 释放座位，离开酒吧
	}
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar2, 10)
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat2(seatId)
	}

	consumers := make(chan int)
	for i := 0; i < cap(bar24x7); i++ {
		go bar24x7.TenServeCustomerAtSeat(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}
