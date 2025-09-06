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
	for i, t := 0, r.Int(); i < t; i++ {
		n, m := r.Int(), r.Int()
		a, asum, b, bsum := make([]int, n), 0, make([]int, n), 0
		for i := 0; i < n; i++ {
			a[i] = r.Int()
			asum += a[i]
		}
		for i := 0; i < n; i++ {
			b[i] = r.Int()
			bsum += b[i]
		}
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
		sort.Sort(sort.IntSlice(b))

		c, j := 0, 0
		for i := 0; i < n; i++ {
			for j < n && a[i]+b[j] < m {
				j++
			}
			if j < n {
				c++
				j++
			}
		}

		w.Println(asum + bsum - m*c)
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
func (r *Reader) Int() int       { return Atoi(r.String()) }
func (r *Reader) String() string { r.sc.Scan(); return r.sc.Text() }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Println(a ...any)  { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()            { w.bf.Flush() }

type Cmp[T Ord] func(s []T, x T) func(int) int
