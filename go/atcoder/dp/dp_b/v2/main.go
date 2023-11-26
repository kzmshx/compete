package main

import (
	"fmt"
	"math"

	. "golang.org/x/exp/constraints"
)

func main() {
	N, K := read[int](), read[int]()
	H := readSlice[int](N)

	// Pull-based DP
	dp := slice[int](N, math.MaxInt32)
	dp[0] = 0
	for i := 1; i < N; i++ {
		for k := 1; k <= K; k++ {
			if i-k >= 0 {
				dp[i] = min(dp[i], dp[i-k]+abs(H[i]-H[i-k]))
			}
		}
	}

	fmt.Println(dp[N-1])
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

// slice returns a slice of length n with each element set to v.
func slice[T any](n int, v T) []T {
	r := make([]T, n)
	for i := 0; i < n; i++ {
		r[i] = v
	}
	return r
}

// min returns the minimum value of a and b.
func min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// abs returns the absolute value of x.
func abs[T Integer | Float](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}
