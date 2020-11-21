package main

import "fmt"

func main() {
	for {
		a := 2
		var b int
		fmt.Printf("请输入猜想的数字:")
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
