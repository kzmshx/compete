package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Reader, w *Writer) {
	for t := r.Int(); t > 0; t-- {
		n, s := r.Int(), r.Str()
		dp := make([]bool, 1<<n)
		for i := 0; i < n; i++ {
			state := (1 << i)
			if s[state-1] == '0' {
				dp[state] = true
			}
		}
		for mask := 1; mask < (1 << n); mask++ {
			if !dp[mask] {
				continue
			}
			for i := 0; i < n; i++ {
				newMask := (mask | (1 << i))
				if s[newMask-1] == '0' {
					dp[newMask] = true
				}
			}
		}
		w.Bool(dp[(1<<n)-1])
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
func (r *Reader) Int() int    { return Atoi(r.Str()) }
func (r *Reader) Str() string { r.sc.Scan(); return r.sc.Text() }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Println(a ...any)  { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Bool(a bool) {
	if a {
		w.Println("Yes")
	} else {
		w.Println("No")
	}
}
func (w *Writer) Flush() { w.bf.Flush() }
