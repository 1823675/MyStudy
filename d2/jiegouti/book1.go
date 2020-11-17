package main

import "fmt"

type books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	fmt.Println(books{"go language", "www.runoob.com", "go 教程", 6407})

	fmt.Println(books{title: "go language", author: "www.runoob.com", subject: "go 教程", book_id: 6407})

	fmt.Println(books{title: "go language", author: "www.runoob.com"})

}
