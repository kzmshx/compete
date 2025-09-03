package main

import (
	"fmt"
)

func main() {
	N := read[int]()
	abc := sliceFunc(N, func(i int) []int { return readSlice[int](3) })

	dp := sliceFunc(N, func(i int) []int { return sliceFill(3, 0) })
	dp[0] = abc[0]
	for i := 1; i < N; i++ {
		for a := 0; a < len(abc[i]); a++ {
			for b := 0; b < len(abc[i]); b++ {
				if a != b {
					chmax(&dp[i][b], dp[i-1][a]+abc[i][b])
				}
			}
		}
	}

	fmt.Println(sliceMax(dp[N-1]))
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

// ordered is a constraint that permits any ordered type: any type
type ordered interface {
	integer | float | ~string
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
