package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	n := r.Int()
	coords := make([][2]int, n)
	for i := 0; i < n; i++ {
		coords[i] = [2]int{r.Int(), r.Int()}
	}
	ans := BinarySearch(0, Pow(10, 9), func(i int) bool {
		cur, ok := [2][2]int{{1, Pow(10, 9)}, {1, Pow(10, 9)}}, true
		for _, c := range coords {
			cur, ok = Intersect2D(cur, [2][2]int{
				{Max(1, c[0]-i), Min(Pow(10, 9), c[0]+i)},
				{Max(1, c[1]-i), Min(Pow(10, 9), c[1]+i)},
			})
			if !ok {
				return false
			}
		}
		return ok
	})
	w.Println(ans)
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

func (s *Scanner) scan() bool { return s.sc.Scan() }

func (s *Scanner) text() string { return s.sc.Text() }

func (s *Scanner) String() string { s.scan(); return s.text() }

func (s *Scanner) Int() int { return Atoi(s.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }

func (w *Writer) Println(a ...any) { fmt.Fprintln(w.bf, a...) }

func (w *Writer) Flush() { w.bf.Flush() }

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

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Pow[T Actual](x T, n int) T {
	y := T(1)
	for n > 0 {
		if n&1 == 1 {
			y *= x
		}
		x *= x
		n >>= 1
	}
	return y
}

func BinarySearch[T Integer](l, r T, f func(T) bool) T {
	for l < r {
		m := T(uint(l+r) >> 1)
		if f(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
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

func Intersect1D[T Ordered](a, b [2]T) ([2]T, bool) {
	min, max := Max(a[0], b[0]), Min(a[1], b[1])
	return [2]T{min, max}, min <= max
}

func Intersect2D[T Ordered](a, b [2][2]T) ([2][2]T, bool) {
	rowRange, okRowRange := Intersect1D(a[0], b[0])
	colRange, okColRange := Intersect1D(a[1], b[1])
	return [2][2]T{rowRange, colRange}, okRowRange && okColRange
}
