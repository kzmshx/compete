package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n)
	fmt.Scan(&x)

	answer := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)

		if a <= x {
			answer += a
		}
	}

	fmt.Println(answer)
}
