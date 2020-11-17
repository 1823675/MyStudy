package main

import "fmt"

var a int = 20

func main() {
	var a, b, c int

	a = 10
	b = 20
	c = 0

	fmt.Printf("main中 a = %d\n", a)
	c = sum(a, b)
	fmt.Printf("main中 c = %d\n", c)
}

func sum(a, b int) int {
	fmt.Printf("sum中 a = %d\n", a)
	fmt.Printf("sum中 b = %d\n", b)

	return a + b
}
