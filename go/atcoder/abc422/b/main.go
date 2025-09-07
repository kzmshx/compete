package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Reader, w *Writer) {
	ym, xm := r.Int(), r.Int()
	f := r.Strs(ym)
	for y := 0; y < ym; y++ {
		for x := 0; x < xm; x++ {
			if f[y][x] == '.' {
				continue
			}

			sum := 0
			if 0 < y && f[y-1][x] == '#' {
				sum++
			}
			if y < ym-1 && f[y+1][x] == '#' {
				sum++
			}
			if 0 < x && f[y][x-1] == '#' {
				sum++
			}
			if x < xm-1 && f[y][x+1] == '#' {
				sum++
			}
			if sum != 2 && sum != 4 {
				w.Println("No")
				return
			}
		}
	}
	w.Println("Yes")
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
func (r *Reader) Int() int            { return Atoi(r.Str()) }
func (r *Reader) Str() string         { r.sc.Scan(); return r.sc.Text() }
func (r *Reader) Strs(n int) []string { return SliceFn(n, func(i int) string { return r.Str() }) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Println(a ...any)  { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()            { w.bf.Flush() }

func SliceFn[T any](n int, f func(i int) T) []T {
	a := make([]T, n)
	for i := range a {
		a[i] = f(i)
	}
	return a
}
