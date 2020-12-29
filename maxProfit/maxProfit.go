package main

import "fmt"

func maxProfit(arr []int) (int, int) {
	comprar := arr[0]
	vender := 0
	//revisar el menor

	for x := 0; x < len(arr)-1; x++ {
		if arr[x] < comprar {
			comprar = x
		}
		if arr[x] > vender {
			vender = x
		}
	}
	return comprar, vender

}

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	a, b := maxProfit(arr)
	fmt.Println(a, b)

}
