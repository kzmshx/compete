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
	fmt.Println("Hello, world!")
}
