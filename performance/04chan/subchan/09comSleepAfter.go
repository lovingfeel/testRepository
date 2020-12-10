package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
采用Sleep来实现间隔的时候，如果cancel调用后任务协程正好处于sleep过程中，这时任务是无法被及时取消的。

小结 如果你需要任务能够被及时的取消，那你应该采用time.After来控制调用的间隔。
 */

func main() {
	TestTimerTask1()
	//TestTimerTask2()
}

var wg sync.WaitGroup
func cancelTaskAfter(interval time.Duration, cancel context.CancelFunc) {
	go func(cancel context.CancelFunc) {
		time.Sleep(interval)
		fmt.Println("Cancell task", time.Now())
		cancel()
		wg.Done()
	}(cancel)
}

func TestTimerTask1() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	cancelTaskAfter(time.Second*2, cancel)
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled", time.Now())
				wg.Done()
				return
			default:
				time.Sleep(time.Second * 10)
				fmt.Println("Invoked", time.Now())
			}
		}
	}(ctx)
	wg.Wait()
}

func TestTimerTask2() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	cancelTaskAfter(time.Second*2, cancel)
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled", time.Now())
				wg.Done()
				return
			case <-time.After(time.Second * 10):
				fmt.Println("Invoked", time.Now())
			}
		}
	}(ctx)
	wg.Wait()
}