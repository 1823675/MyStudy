package main

import "fmt"

func main() {
	var numbers []int
	printslice(numbers)

	numbers = append(numbers, 0)
	printslice(numbers)

	numbers = append(numbers, 1)
	printslice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printslice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	copy(numbers1, numbers)
	printslice(numbers1)
}

func printslice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
