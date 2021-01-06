package main

import (
	"fmt"
	"net/http"
)

func wangye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello,world")
}

func main() {
	http.HandleFunc("/", wangye)
	http.ListenAndServe(":8000", nil)
}
