package main

import "fmt"

func main() {
	var ccmap map[string]string
	ccmap = make(map[string]string)

	ccmap["france"] = "巴黎"
	ccmap["italy"] = "罗马"
	ccmap["japan"] = "东京"
	ccmap["india"] = "新德里"

	for country := range ccmap {
		fmt.Println(country, "首都是", ccmap[country])
	}
	capital, ok := ccmap["american"]
	if ok {
		fmt.Println("american的首都是", capital)
	} else {
		fmt.Println("american的首都不存在")
	}
}
