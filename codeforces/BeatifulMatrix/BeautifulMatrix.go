package main

import (
	"fmt"
	"math"
)

func main() {
	var i int

	for x := 1; x <= 5; x++ {

		for z := 1; z <= 5; z++ {
			fmt.Scan(&i)
			if i == 1 {
				fmt.Println()
				fmt.Println(math.Abs(float64(x-3)) + math.Abs(float64(z-3)))
				break
			}
		}

	}
}
