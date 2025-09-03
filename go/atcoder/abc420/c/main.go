package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Reader, w *Writer) {
	n, q := r.Int(), r.Int()
	a, b, sum := r.Ints(n), make([]int, n), 0
	for i := 0; i < n; i++ {
		b[i] = r.Int()
		sum += Min(a[i], b[i])
	}
	for i := 0; i < q; i++ {
		if c, x, v := r.String(), r.Int()-1, r.Int(); c == "A" {
			sum, a[x] = sum-Min(a[x], b[x])+Min(v, b[x]), v
		} else if c == "B" {
			sum, b[x] = sum-Min(a[x], b[x])+Min(a[x], v), v
		} else {
			panic("unexpected query")
		}
		w.Println(sum)
	}
}

func main() {
	r, w := NewReader(os.Stdin, 1*1024*1024), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

type Reader struct{ sc *bufio.Scanner }

func NewReader(r io.Reader, size int) *Reader {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Reader{sc}
}
func (r *Reader) Int() int { return Atoi(r.String()) }
func (r *Reader) Ints(n int) []int {
	return MakeSlice(n, func(i int) int { return r.Int() })
}
func (r *Reader) String() string { r.sc.Scan(); return r.sc.Text() }

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

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MakeSlice[T any](n int, f func(i int) T) []T {
	a := make([]T, n)
	for i := range a {
		a[i] = f(i)
	}
	return a
}

type CyclicInt[T Signed] struct {
	v T
	c T
}

type ModInt[T Integer] struct {
	v T
	m T
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
