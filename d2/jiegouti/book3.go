package main

import "fmt"

type books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book1 books
	var book2 books

	book1.title = "go language"
	book1.author = "www.runoob.com"
	book1.subject = "go teach"
	book1.book_id = 6407

	book2.title = "py language"
	book2.author = "www.runoob.com"
	book2.subject = "py teach"
	book2.book_id = 6400

	printbook(&book1)

	printbook(&book2)
}
func printbook(book *books) {
	fmt.Printf("book  title : %s\n", book.title)
	fmt.Printf("book  author : %s\n", book.author)
	fmt.Printf("book  subject : %s\n", book.subject)
	fmt.Printf("book  book_id : %d\n", book.book_id)
}
