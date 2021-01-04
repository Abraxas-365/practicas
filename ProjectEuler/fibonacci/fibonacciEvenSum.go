package main

import "fmt"

func fibo(n int) int {
	a := 0
	b := 2
	t := 2
	s := 0

	for t <= n {
		s += t
		t = a + 4*b
		a = b
		b = t
	}
	return s
}

func main() {

	fmt.Println(fibo(4e6))
}
