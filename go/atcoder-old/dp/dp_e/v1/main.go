package main

import (
	"fmt"
	"math"
)

func main() {
	N, W := read[int](), read[int]()
	ws, vs, vSum := make([]int, N), make([]int, N), 0
	for i := 0; i < N; i++ {
		ws[i], vs[i] = read[int](), read[int]()
		vSum += vs[i]
	}

	// dp[i][v]: i番目までの品物を使って、価値がv以上にするための、重さの最小値
	// 最小化問題なので、初期値は十分大きな値にしておく
	// 1次元目のインデックスで、0番目からi番目までの品物を使うことを表現する
	// 2次元目のインデックスで、価値が0からvSum（vの合計値）までの範囲を表現する
	dp := sliceFunc(N+1, func(i int) []int { return sliceFill(vSum+1, math.MaxInt32) })
	dp[0][0] = 0
	for i := 0; i < N; i++ {
		wi, vi := ws[i], vs[i]
		for v := 0; v <= vSum; v++ {
			chmin(&dp[i+1][v], dp[i][v])
			if v+vi <= vSum {
				chmin(&dp[i+1][v+vi], dp[i][v]+wi)
			}
		}
	}

	for i := vSum; i >= 0; i-- {
		if dp[N][i] <= W {
			fmt.Println(i)
			return
		}
	}
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

// chmin sets the minimum value of a and b to a and returns the minimum value.
func chmin[T ordered](a *T, b T) T {
	if *a > b {
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
