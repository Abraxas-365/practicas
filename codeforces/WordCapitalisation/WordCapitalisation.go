package main

import "fmt"

func main() {
	var str string
	var nuevo string
	fmt.Scan(&str)
	if str[0] > 90 {
		nuevo = string(str[0] - 32)
	} else {
		nuevo = string(str[0])
	}
	for x := 1; x < len(str); x++ {
		nuevo = nuevo + string(str[x])
	}
	fmt.Println(nuevo)
}
