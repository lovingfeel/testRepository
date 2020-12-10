package main

import (
	"fmt"
	"time"
)

func main()  {
	//rand.Seed(time.Now().UnixNano())

	if n := 1; n%2 == 0 {
		fmt.Println(n, "是一个偶数。")
	} else {
		fmt.Println(n, "是一个奇数。")
	}

	b := 4 // 此n不是上面声明的n
	if b % 2 == 0 {
		fmt.Println("一个偶数。")
	}

	if  3 % 2 != 0 {
		fmt.Println("一个奇数。")
	}

	if h := time.Now().Hour(); h < 12 {
		fmt.Println("现在为上午。")
	} else if h > 19 {
		fmt.Println("现在为晚上。")
	} else {
		fmt.Println("现在为下午。")
		// 左h是一个新声明的变量，右h已经在上面声明了。
		h := h
		// 刚声明的h遮掩了上面声明的h。
		_ = h
	}
}
