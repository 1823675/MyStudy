package main

import "fmt"

func main() {
	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("a dizhi:%x\n", &a)

	fmt.Printf("ip zhizhen chucun dizhi:%x\n", ip)

	fmt.Printf("*ip mianliang zhi:%d\n", *ip)
}
