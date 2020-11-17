package main

import "fmt"

func main() {
	var numbers []int

	printslice(numbers)

	if numbers == nil {
		fmt.Printf("qiepian shi kongde ")
	}
}
func printslice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
