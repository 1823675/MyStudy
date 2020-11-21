package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	a := rand.Int() % 10
	for {
		var b int
		fmt.Println("请输入猜想的数字:")
		fmt.Scan(&b)
		if a > b {
			fmt.Println("猜小了")
		} else if a < b {
			fmt.Println("猜大了")
		} else {
			fmt.Println("猜对了")
			break
		}
	}
}
