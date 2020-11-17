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

	fmt.Printf("book 1 title : %s\n", book1.title)
	fmt.Printf("book 1  author: %s\n", book1.author)
	fmt.Printf("book 1  subject : %s\n", book1.subject)
	fmt.Printf("book 1  book_id : %d\n", book1.book_id)

	fmt.Printf("book 2 title : %s\n", book2.title)
	fmt.Printf("book 2 author : %s\n", book2.author)
	fmt.Printf("book 2 subject : %s\n", book2.subject)
	fmt.Printf("book 2 book_id : %d\n", book2.book_id)
}
