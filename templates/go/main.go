package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}

// signed is a constraint that permits any signed integer type.
type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// unsigned is a constraint that permits any unsigned integer type.
type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// integer is a constraint that permits any integer type.
type integer interface {
	signed | unsigned
}

// float is a constraint that permits any floating-point type.
type float interface {
	~float32 | ~float64
}

// actual is a constraint that permits any complex numeric type.
type actual interface {
	integer | float
}

// imaginary is a constraint that permits any complex numeric type.
type imaginary interface {
	~complex64 | ~complex128
}

// ordered is a constraint that permits any ordered type: any type
type ordered interface {
	integer | float | ~string
}

// addable is a constraint that permits any ordered type: any type
type addable interface {
	integer | float | imaginary | ~string
}

// read reads a value from stdin.
func read[T any]() (r T) {
	fmt.Scan(&r)
	return r
}

// readSlice reads n values from stdin.
func readSlice[T any](n int) []T {
	r := make([]T, n)
	for i := 0; i < n; i++ {
		r[i] = read[T]()
	}
	return r
}

// chmax sets the maximum value of a and b to a and returns the maximum value.
func chmax[T ordered](a *T, b T) T {
	if *a < b {
		*a = b
	}
	return *a
}

// chmin sets the minimum value of a and b to a and returns the minimum value.
func chmin[T ordered](a *T, b T) T {
	if *a > b {
		*a = b
	}
	return *a
}

// max returns the maximum value of a and b.
func max[T ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// min returns the minimum value of a and b.
func min[T ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// abs returns the absolute value of x.
func abs[T actual](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

// sliceFill returns a slice of length n with each element set to v.
func sliceFill[T any](n int, v T) []T {
	r := make([]T, n)
	for i := 0; i < n; i++ {
		r[i] = v
	}
	return r
}

// sliceFunc returns a slice of length n with each element set to the result of f(i).
func sliceFunc[T any](n int, f func(int) T) []T {
	r := make([]T, n)
	for i := 0; i < n; i++ {
		r[i] = f(i)
	}
	return r
}

// sliceMax returns the maximum value of s.
func sliceMax[T ordered](s []T) (r T) {
	for _, v := range s {
		chmax(&r, v)
	}
	return r
}

// sliceSum returns the sliceSum of values.
func sliceSum[T addable](s []T) (r T) {
	for _, v := range s {
		r += v
	}
	return r
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
