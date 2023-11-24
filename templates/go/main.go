package main

import (
	"fmt"

	. "golang.org/x/exp/constraints"
)

// addable is the type of values that support addition.
type addable interface {
	Integer | Float | Complex | string
}

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

// sum returns the sum of values.
func sum[T addable](values []T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
}

// gcd returns the greatest common divisor of a and b.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm returns the least common multiple of a and b.
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// min returns the minimum value of a and b.
func min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum value of a and b.
func max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// minmax returns the minimum and maximum values of a and b.
func minmax[T Ordered](a, b T) (T, T) {
	if a < b {
		return a, b
	}
	return b, a
}

// abs returns the absolute value of x.
func abs[T Integer | Float](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

func main() {
	fmt.Println("Hello, world!")
}
