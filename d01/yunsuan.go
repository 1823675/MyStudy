package main 
import "fmt"
func main()  {
	var a int = 21
	var b int =10
	var c int = a + b 
	fmt.Printf("c值为%d\n",c)
	c = a - b 
	fmt.Printf("c值为%d\n",c)
	c = a * b 
	fmt.Printf("c值为%d\n",c)
	c = a / b 
	fmt.Printf("c值为%d\n",c)
	c = a % b 
	fmt.Printf("c值为%d\n",c)
	a++
	fmt.Printf("a值为%d\n",a)
	a = 21
	a--
	fmt.Printf("a值为%d\n",a)
}