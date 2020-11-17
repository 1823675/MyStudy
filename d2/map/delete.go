package main

import "fmt"

func main() {
	ccmap := map[string]string{"france": "paris", "italy": "rome", "japan": "tokyo", "india": "new delhi"}

	fmt.Println("yuanshimap")

	for country := range ccmap {
		fmt.Println(country, "capitaol is ", ccmap[country])
	}
	fmt.Println("france is delete")

	fmt.Println("delete map")

	for country := range ccmap {
		fmt.Println(country, "captial is", ccmap[country])
	}
}
