package main

import (
	"fmt"
)

func main() {
	s, t := read[string](), read[string]()

	dp := sliceFunc(len(s)+1, func(i int) []int { return sliceFill(len(t)+1, 0) })
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				chmax(&dp[i+1][j+1], dp[i][j]+1)
			}
			chmax(&dp[i+1][j+1], dp[i+1][j])
			chmax(&dp[i+1][j+1], dp[i][j+1])
		}
	}

	lcs := ""
	i, j := len(s), len(t)
	for i > 0 && j > 0 {
		if dp[i][j] == dp[i-1][j] {
			i--
		} else if dp[i][j] == dp[i][j-1] {
			j--
		} else {
			lcs = string(s[i-1]) + lcs
			i--
			j--
		}
	}

	fmt.Println(lcs)
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
