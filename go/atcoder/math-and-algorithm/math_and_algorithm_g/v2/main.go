package main

import (
	"fmt"
)

func Input[T any]() T {
	var value T
	fmt.Scan(&value)
	return value
}

func Inputs[T any](n int) []T {
	values := make([]T, n)
	for i := 0; i < n; i++ {
		values[i] = Input[T]()
	}
	return values
}

type Addable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		complex64 | complex128 |
		string
}

func Sum[T Addable](values []T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
}

func main() {
	n, x, y := Input[int](), Input[int](), Input[int]()
	m := map[int]struct{}{}
	for t := x; t <= n; t += x {
		m[t] = struct{}{}
	}
	for t := y; t <= n; t += y {
		m[t] = struct{}{}
	}
	fmt.Println(len(m))
}
