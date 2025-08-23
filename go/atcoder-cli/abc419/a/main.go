package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solve(r *Scanner, w *Writer) {
	switch r.String() {
	case "red":
		w.Println("SSS")
	case "blue":
		w.Println("FFF")
	case "green":
		w.Println("MMM")
	default:
		w.Println("Unknown")
	}
}

func main() {
	r, w := NewScanner(os.Stdin, MaxBufferSize), NewWriter(os.Stdout)
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

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }

func (w *Writer) Println(a ...any) { fmt.Fprintln(w.bf, a...) }

func (w *Writer) Flush() { w.bf.Flush() }

type UnionFind struct {
	parent []int
	size   []int
}
