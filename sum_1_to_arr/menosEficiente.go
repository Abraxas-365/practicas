package main

import "fmt"

func sum(arr []int) []int {
	suma := 0
	for x := 0; x < len(arr); x++ {
		suma = suma + arr[x]
		suma = suma * 10

	}
	suma = suma / 10

	suma = suma + 1
	nuevo := []int{}
	for suma != 0 {
		i := suma % 10
		nuevo = append([]int{i}, nuevo...)
		suma = suma / 10
	}
	return nuevo
}

func main() {
	arr := []int{9, 9, 9}
	fmt.Println(sum(arr))
}
