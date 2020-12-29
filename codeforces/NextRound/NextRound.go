package main

import "fmt"

func main() {

	var participantes int
	var mayor int
	fmt.Scan(&participantes)
	fmt.Scan(&mayor)
	count := 0

	arr := make([]int, participantes)

	for i := 0; i < participantes; i++ {

		if _, err := fmt.Scan(&arr[i]); err != nil {
			panic(err)
		}
	}
	for _, val := range arr {
		if val >= arr[mayor] {
			count++
		}
	}

	fmt.Println(count)
}
