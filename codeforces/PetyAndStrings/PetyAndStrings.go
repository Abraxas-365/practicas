package main

import (
	"fmt"

	"strings"
)

func main() {
	var (
		primera string
		segunda string
	)

	fmt.Scan(&primera)

	fmt.Scan(&segunda)
	primera = strings.ToLower(primera)
	segunda = strings.ToLower(segunda)
	count1 := 0
	count2 := 0
	for _, x := range primera {
		count1 = count1 + int(x)
	}

	for _, x := range segunda {
		count2 = count2 + int(x)
	}

	if primera == segunda {
		fmt.Println(0)
	} else if primera > segunda {
		fmt.Println(1)

	} else if primera < segunda {
		fmt.Println(-1)
	}

}
