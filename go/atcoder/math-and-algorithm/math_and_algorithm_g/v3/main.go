package main

import (
	"fmt"
)

// Input returns a value.
func Input[T any]() T {
	var value T
	fmt.Scan(&value)
	return value
}

// Inputs returns n values.
func Inputs[T any](n int) []T {
	values := make([]T, n)
	for i := 0; i < n; i++ {
		values[i] = Input[T]()
	}
	return values
}

// Addable is the type of values that support addition.
type Addable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		complex64 | complex128 |
		string
}

// Sum returns the sum of values.
func Sum[T Addable](values []T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
}

// GCD returns the greatest common divisor of a and b.
func GCD(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of a and b.
func LCM(a int, b int) int {
	return a * b / GCD(a, b)
}

func main() {
	n, x, y := Input[int](), Input[int](), Input[int]()
	fmt.Println(n/x + n/y - n/LCM(x, y))
}
