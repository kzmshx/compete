package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solve(r *Reader, w *Writer) {
	s := r.Bytes()
	var ts []int
	for i, c := range s {
		if c == 't' {
			ts = append(ts, i)
		}
	}
	ans := .0
	for i := 0; i < len(ts)-1; i++ {
		for j := i + 1; j < len(ts); j++ {
			l, r := ts[i], ts[j]
			if r-l >= 2 {
				ChooseMax(&ans, float64(j-i-1)/float64(r-l-1))
			}
		}
	}
	w.PrintlnFloat64(ans, 10)
}

func main() {
	r, w := NewReader(os.Stdin, 1*1024*1024), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

type Signed interface{ ~int | ~int32 | ~int64 }
type Unsigned interface{ ~uint | ~uint32 | ~uint64 }
type Int interface{ Signed | Unsigned }
type Float interface{ ~float32 | ~float64 }
type Num interface{ Int | Float }
type Ord interface{ Int | Float | ~string }

type Reader struct{ sc *bufio.Scanner }

func NewReader(r io.Reader, size int) *Reader {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Reader{sc}
}
func (r *Reader) Bytes() []byte { r.sc.Scan(); return r.sc.Bytes() }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer                  { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Printf(format string, a ...any)     { fmt.Fprintf(w.bf, format, a...) }
func (w *Writer) PrintlnFloat64(a float64, prec int) { w.Printf("%.*f", prec, a) }
func (w *Writer) Flush()                             { w.bf.Flush() }

func ChooseMax[T Ord](a *T, b T) T {
	if *a < b {
		*a = b
	}
	return *a
}
