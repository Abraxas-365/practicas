package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	if i > 100 {
		fmt.Println("NO")
	}
	if i%2 == 0 && i > 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
