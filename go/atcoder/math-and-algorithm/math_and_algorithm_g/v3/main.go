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

// gcd returns the greatest common divisor of a and b.
func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm returns the least common multiple of a and b.
func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	n, x, y := input[int](), input[int](), input[int]()
	fmt.Println(n/x + n/y - n/lcm(x, y))
}
