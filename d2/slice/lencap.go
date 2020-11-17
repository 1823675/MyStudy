package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)

	printslice(numbers)
}
func printslice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
