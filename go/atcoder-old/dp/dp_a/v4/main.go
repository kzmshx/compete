package main

import (
	"fmt"
	"math"

	. "golang.org/x/exp/constraints"
)

func costRec(H []int, i int, memo []int) int {
	if memo[i] != math.MaxInt32 {
		return memo[i]
	} else if i == 0 {
		return chmin(&memo[i], 0)
	} else if i == 1 {
		return chmin(&memo[i], costRec(H, i-1, memo)+abs(H[i]-H[i-1]))
	} else {
		return chmin(&memo[i], min(costRec(H, i-1, memo)+abs(H[i]-H[i-1]), costRec(H, i-2, memo)+abs(H[i]-H[i-2])))
	}
}

func cost(H []int, i int) int {
	return costRec(H, i, slice[int](len(H), math.MaxInt32))
}

func main() {
	N := read[int]()
	H := readSlice[int](N)
	fmt.Println(cost(H, N-1))
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

// chmin updates a to the minimum of a and b.
func chmin[T Ordered](a *T, b T) T {
	if *a > b {
		*a = b
	}
	return *a
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
