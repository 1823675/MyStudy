package main
import "fmt"
func main()  {
	var a int = 20
	var ip *int
	fmt.Printf("ip var = %d\n", ip )
	ip = &a 
	fmt.Printf("a var address: %x\n", &a )
	fmt.Printf("ip var address: %x\n", ip )
	fmt.Printf("*ip var = %d\n", *ip )
}