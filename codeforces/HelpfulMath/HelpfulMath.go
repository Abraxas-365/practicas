package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{}

	var str string

	fmt.Scan(&str)

	for _, x := range str {
		if x == '1' || x == '2' || x == '3' {

			arr = append(arr, string(x))
		}
	}

	sort.Strings(arr)
	fmt.Print(arr[0])
	for x := 1; x < len(arr); x++ {

		fmt.Print("+", arr[x])
	}

}
