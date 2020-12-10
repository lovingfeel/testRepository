package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)
/**
通知可以被看作是特殊的请求/回应用例。在一个通知用例中，我们并不关心回应的值，
我们只关心回应是否已发生。 所以我们常常使用空结构体类型struct{}来做为通道的元素类型
，因为空结构体类型的尺寸为零，能够节省一些内存（虽然常常很少量）。
 */
func main() {
	values := make([]byte, 32 * 1024 * 1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	done := make(chan struct{}) // 也可以是缓冲的

	// 排序协程
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // 通知排序已完成
	}()

	// 并发地做一些其它事情...

	<- done // 等待通知
	fmt.Println(values[0], values[len(values)-1])
}
