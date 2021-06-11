package main
import "fmt"
type Books struct {
	title string
	author string
	subject string 
	book_id int 
}
func main()  {
	fmt.Println(Books{"go language", "www.runoob.com", "go teach", 6495407})
	fmt.Println(Books{title: "go language", author: "www.runoob.com" , subject: "go teach", book_id: 6495407})
	fmt.Println(Books{title: "go language", author: "www.runoob.com"})
}