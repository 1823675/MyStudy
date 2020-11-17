package main

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("第1行a 变量类型为 = %T\n", a)
	fmt.Printf("第2行b 变量类型为 = %T\n", b)
	fmt.Printf("第3行c 变量类型为 = %T\n", c)

	ptr = &a
	fmt.Printf("a 的值为 = %d\n", a)
	fmt.Printf("*ptr 的值为 = %d\n", *ptr)
}
