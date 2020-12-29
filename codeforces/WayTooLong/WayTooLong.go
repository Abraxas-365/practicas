package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	var i int
	es := []string{}
	fmt.Scan(&i)
	for x := 0; x < i; x++ {
		fmt.Scan(&s)

		if len(s) > 10 {
			es = append(es, string(s[0])+strconv.Itoa(len(s)-2)+string(s[len(s)-1]))
		} else {
			es = append(es, s)
		}

	}

	for _, val := range es {
		fmt.Println(val)
	}
}
