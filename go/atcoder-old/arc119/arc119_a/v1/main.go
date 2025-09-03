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
	N := r.Int()
	ans := math.MaxInt64
	for b := 0; Pow(2, b) <= N; b++ {
		d := Pow(2, b)
		ChooseMin(&ans, b+N/d+N%d)
	}
	w.Println(ans)
}

func main() {
	r, w := NewScanner(os.Stdin, maxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

type signed interface{ ~int | ~int32 | ~int64 }

type unsigned interface{ ~uint | ~uint32 | ~uint64 }

type integer interface{ signed | unsigned }

type float interface{ ~float32 | ~float64 }

type actual interface{ integer | float }

type ordered interface{ integer | float | ~string }

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

func ChooseMin[T ordered](a *T, b T) T {
	if *a > b {
		*a = b
	}
	return *a
}

func Pow[T actual](x T, n int) T {
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

const maxBufferSize = 1 * 1024 * 1024

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

func (w *Writer) Println(a ...interface{}) { fmt.Fprintln(w.bf, a...) }

func (w *Writer) Flush() { w.bf.Flush() }

type UnionFind struct {
	parent []int
	size   []int
}
