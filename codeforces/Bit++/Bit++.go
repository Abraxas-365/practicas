package main

import "fmt"

func main() {
	sum := 0
	var cant int
	fmt.Scan(&cant)

	var operacion string
	for x := 0; x < cant; x++ {
		fmt.Scan(&operacion)
		if string(operacion[0]) == "+" && string(operacion[1]) == "+" || string(operacion[len(operacion)-1]) == "+" && string(operacion[len(operacion)-1]) == "+" {
			sum++
		} else if operacion[0] == '-' && operacion[1] == '-' || operacion[len(operacion)-1] == '-' && operacion[len(operacion)-1] == '-' {
			sum--
		}

	}
	fmt.Println(sum)
}
