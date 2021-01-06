package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	data := rand.Int63n(20)
	fmt.Println(data)

}
