package main

import "fmt"

func main() {
	a := 2
	var b int
	fmt.Printf("请输入猜想的数字:")
	fmt.Scan(&b)
	if a > b {
		fmt.Printf("猜小了")
	} else if a < b {
		fmt.Printf("猜大了")
	} else {
		fmt.Printf("猜对了")
	}
}
