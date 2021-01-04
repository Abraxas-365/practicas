package main

import "fmt"

func primefactors(n int) {

	d := 2
	for n > 1 {

		for n%d == 0 {

			n /= d
			fmt.Println(d)

		}
		d = d + 1
	}

}
func main() {
	primefactors(600851475143)
}
