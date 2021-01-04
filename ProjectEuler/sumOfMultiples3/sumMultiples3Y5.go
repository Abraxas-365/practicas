package main

import "fmt"

func main() {
	var hasta int
	sum := 0
	fmt.Scan(&hasta)
	for x := 0; x < hasta; x++ {
		if x%3 == 0 || x%5 == 0 {

			sum = sum + x
		}
	}
	fmt.Println(sum)
}
