package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	//这一行是哪里不对呀
	fmt.Println("%v %v %v %q\n", i, f, b, s)
}
