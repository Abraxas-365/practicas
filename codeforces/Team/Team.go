package main

import "fmt"

func main() {
	var cant int
	fmt.Scan(&cant)
	num := 0

	var a int
	var b int
	var c int

	for x := 0; x < cant; x++ {
		fmt.Scan(&a, &b, &c)
		if a+b+c >= 2 {
			num = num + 1
		}
	}
	fmt.Print(num)
}
