package main

import (
	"fmt"
)

// input returns a value.
func input[T any]() T {
	var value T
	fmt.Scan(&value)
	return value
}

func main() {
	n, s := input[int](), input[int]()
	c := 0
	for r := 1; r <= n; r++ {
		for b := 1; b <= n; b++ {
			if r+b <= s {
				c++
			}
		}
	}
	fmt.Println(c)
}
