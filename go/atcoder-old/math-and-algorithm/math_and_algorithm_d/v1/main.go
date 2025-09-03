package main

import (
	"fmt"
)

func Scan[T any]() T {
	var value T
	fmt.Scan(&value)
	return value
}

func main() {
	a1, a2, a3 := Scan[int](), Scan[int](), Scan[int]()
	fmt.Println(a1 * a2 * a3)
}
