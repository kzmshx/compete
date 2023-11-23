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

// inputs returns n values.
func inputs[T any](n int) []T {
	values := make([]T, n)
	for i := 0; i < n; i++ {
		values[i] = input[T]()
	}
	return values
}

// addable is the type of values that support addition.
type addable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		complex64 | complex128 |
		string
}

// sum returns the sum of values.
func sum[T addable](values []T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
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
	fmt.Println("Hello, world!")
}
