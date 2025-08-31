package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Reader, w *Writer) {
	n, m := r.Int(), r.Int()
	bs := MakeSlice(m, func(i int) []byte { return make([]byte, n) })
	for i := 0; i < n; i++ {
		s := r.Bytes()
		for j := 0; j < m; j++ {
			bs[j][i] = s[j]
		}
	}

	a := make([]int, n)
	max := 0
	for _, bsi := range bs {
		for j := 0; j < n; j++ {
			c := SliceCount(bsi, func(b byte) bool { return b == '1' })
			if c == 0 || c == n || (2*c < n && bsi[j] == '1') || (2*c > n && bsi[j] == '0') {
				a[j]++
				ChooseMax(&max, a[j])
			}
		}
	}

	for i := 0; i < n; i++ {
		if a[i] == max {
			w.Println(i + 1)
		}
	}
}

func main() {
	r, w := NewReader(os.Stdin, MaxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

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

const MaxBufferSize = 1 * 1024 * 1024

type Reader struct{ sc *bufio.Scanner }

func NewReader(r io.Reader, size int) *Reader {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Reader{sc}
}
func (r *Reader) Bytes() []byte  { r.sc.Scan(); return r.sc.Bytes() }
func (r *Reader) String() string { r.sc.Scan(); return r.sc.Text() }
func (r *Reader) Int() int       { return Atoi(r.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Println(a ...any)  { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()            { w.bf.Flush() }

func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

func ChooseMax[T Ordered](a *T, b T) T {
	if *a < b {
		*a = b
	}
	return *a
}

func MakeSlice[T any](n int, f func(i int) T) []T {
	a := make([]T, n)
	for i := range a {
		a[i] = f(i)
	}
	return a
}

func SliceCount[T any](a []T, f func(T) bool) (count int) {
	for _, v := range a {
		if f(v) {
			count++
		}
	}
	return
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
