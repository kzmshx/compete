package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	n, m, l := r.Int(), r.Int(), r.Int()
	a := ScanSlice(r, n, func(s *Scanner) int { return s.Int() })

	groups := make([][]int, l)
	for i := 0; i < n; i++ {
		groups[i%l] = append(groups[i%l], a[i])
	}

	costs := make([][]int, l)
	for gi := 0; gi < l; gi++ {
		costs[gi] = make([]int, m)
		for t := 0; t < m; t++ {
			cost := 0
			for _, v := range groups[gi] {
				cost += (t - v + m) % m
			}
			costs[gi][t] = cost
		}
	}

	dp := make([][]int, l+1)
	for gi := 0; gi <= l; gi++ {
		dp[gi] = make([]int, m)
		for t := 0; t < m; t++ {
			dp[gi][t] = math.MaxInt
		}
	}
	dp[0][0] = 0

	for gi := 0; gi < l; gi++ {
		for prevSum := 0; prevSum < m; prevSum++ {
			if dp[gi][prevSum] == math.MaxInt {
				continue
			}
			for t := 0; t < m; t++ {
				newSum := (prevSum + t) % m
				newCost := dp[gi][prevSum] + costs[gi][t]
				ChooseMin(&dp[gi+1][newSum], newCost)
			}
		}
	}

	w.Println(dp[l][0])
}

func main() {
	r, w := NewScanner(os.Stdin, MaxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

const MaxBufferSize = 1 * 1024 * 1024

type Scanner struct{ sc *bufio.Scanner }

func NewScanner(r io.Reader, size int) *Scanner {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}
func (s *Scanner) scan() bool     { return s.sc.Scan() }
func (s *Scanner) text() string   { return s.sc.Text() }
func (s *Scanner) String() string { s.scan(); return s.text() }
func (s *Scanner) Int() int       { return Atoi(s.String()) }
func ScanSlice[T any](s *Scanner, n int, f func(s *Scanner) T) []T {
	a := make([]T, n)
	for i := range a {
		a[i] = f(s)
	}
	return a
}

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Println(a ...any)  { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()            { w.bf.Flush() }

type Signed interface{ ~int | ~int32 | ~int64 }
type Unsigned interface{ ~uint | ~uint32 | ~uint64 }
type Integer interface{ Signed | Unsigned }
type Float interface{ ~float32 | ~float64 }
type Actual interface{ Integer | Float }
type Imaginary interface{ ~complex64 | ~complex128 }
type Ordered interface{ Integer | Float | ~string }
type Addable interface {
	Integer | Float | Imaginary | ~string
}

func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

func ChooseMin[T Ordered](a *T, b T) T {
	if *a > b {
		*a = b
	}
	return *a
}

type UnionFind struct {
	parent []int
	size   []int
}

type priorityQueueItem[T any, P Ordered] struct {
	value    T
	priority P
}

type PriorityQueue[T any, P Ordered] struct {
	items      []*priorityQueueItem[T, P]
	itemCount  uint
	comparator func(lhs, rhs P) bool
}

type Diff[T Actual] struct{ delta []T }

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
