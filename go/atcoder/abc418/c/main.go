package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func Solve(r *Reader, w *Writer) {
	n, q := r.Int(), r.Int()
	a := r.Ints(n)
	sort.Ints(a)

	asum := make([]int, n+1)
	asum[0] = 0
	for i, v := range a {
		asum[i+1] = asum[i] + v
	}

	for i := 0; i < q; i++ {
		b := r.Int()
		if j := LowerBound(a, b); j < len(a) {
			w.Println(asum[j] + (b-1)*(len(a)-j) + 1)
		} else {
			w.Println(-1)
		}
	}
}

func main() {
	r, w := NewReader(os.Stdin, 1*1024*1024), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

type Sgn interface{ ~int | ~int32 | ~int64 }
type Uns interface{ ~uint | ~uint32 | ~uint64 }
type Int interface{ Sgn | Uns }
type Float interface{ ~float32 | ~float64 }
type Num interface{ Int | Float }
type Ord interface{ Int | Float | ~string }

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

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

func BinarySearch[T Int](l, r T, f func(T) bool) T {
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

func MakeSlice[T any](n int, f func(i int) T) []T {
	a := make([]T, n)
	for i := range a {
		a[i] = f(i)
	}
	return a
}

func LowerBound[T Ord](s []T, x T) int {
	return BinarySearch(0, len(s), func(i int) bool { return s[i] >= x })
}
