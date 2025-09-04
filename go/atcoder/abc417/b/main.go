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
	a := r.Ints(n)
	for i := 0; i < m; i++ {
		if j := BinarySearchExact(a, r.Int()); j != -1 {
			a = append(a[:j], a[j+1:]...)
		}
	}
	w.PrintlnInts(a)
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
func (w *Writer) Print(a ...any)    { fmt.Fprint(w.bf, a...) }
func (w *Writer) Println(a ...any)  { fmt.Fprintln(w.bf, a...) }
func (w *Writer) PrintlnInts(a []int) {
	for i, v := range a {
		w.Print(v)
		if i < len(a)-1 {
			w.Print(" ")
		}
	}
	w.Println()
}
func (w *Writer) Flush() { w.bf.Flush() }

func MakeSlice[T any](n int, f func(i int) T) []T {
	a := make([]T, n)
	for i := range a {
		a[i] = f(i)
	}
	return a
}

func BinarySearch(l, r int, f func(int) int) (int, bool) {
	for l < r {
		m := l + (r-l)/2
		switch cmp := f(m); {
		case cmp == 0:
			return m, true
		case cmp < 0:
			l = m + 1
		case cmp > 0:
			r = m
		}
	}
	return l, false
}

type Cmp[T Ord] func(s []T, x T) func(int) int

func Cmp3Way[T Ord](s []T, x T) func(int) int {
	return func(i int) int {
		if s[i] == x {
			return 0
		}
		if s[i] < x {
			return -1
		}
		return 1
	}
}

func BinarySearchExact[T Ord](s []T, x T) int {
	if i, ok := BinarySearch(0, len(s), Cmp3Way(s, x)); ok {
		return i
	}
	return -1
}
