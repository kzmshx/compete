package main

import (
	"fmt"
	"math"

	. "golang.org/x/exp/constraints"
)

/**
コストを無限大で初期化しておき緩和していく方針で解く
*/

func main() {
	n := input[int]()
	h := inputs[int](n)

	dp := slice[int](n, math.MaxInt32)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = min(dp[i], dp[i-1]+abs(h[i]-h[i-1]))
		if i > 1 {
			dp[i] = min(dp[i], dp[i-2]+abs(h[i]-h[i-2]))
		}
	}

	fmt.Println(dp[n-1])
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

func slice[T any](n int, value T) []T {
	values := make([]T, n)
	for i := 0; i < n; i++ {
		values[i] = value
	}
	return values
}

// min returns the minimum value of values.
func min[T Ordered](a T, b ...T) T {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}

// abs returns the absolute value of x.
func abs[T Integer | Float](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}
