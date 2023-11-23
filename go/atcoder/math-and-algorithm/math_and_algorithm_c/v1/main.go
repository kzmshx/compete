package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := 0
	for i := 0; i < n; i++ {
		var ai int
		fmt.Scan(&ai)
		a += ai
	}

	fmt.Println(a)
}
