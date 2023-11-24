package main

import (
	"fmt"

	. "golang.org/x/exp/constraints"
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

func main() {
	n := input[int]()
	hs := inputs[int](n)

	dp := make([]int, n)
	dp[0], dp[1] = 0, abs(hs[1]-hs[0])
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1]+abs(hs[i]-hs[i-1]), dp[i-2]+abs(hs[i]-hs[i-2]))
	}

	fmt.Println(dp[n-1])
}
