package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	N := r.Int()
	ans := 0
	for i := 0; i < N; i++ {
		L, R := r.Int(), r.Int()
		ans += R - L + 1
	}
	w.Println(ans)
}

func main() {
	r, w := NewScanner(os.Stdin, maxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

func unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func atoi(s string) int { return unwrap(strconv.Atoi(s)) }

const maxBufferSize = 1 * 1024 * 1024

type Scanner struct{ sc *bufio.Scanner }

func NewScanner(r io.Reader, size int) *Scanner {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}
func (s *Scanner) scan() bool     { return s.sc.Scan() }
func (s *Scanner) text() string   { return s.sc.Text() }
func (s *Scanner) String() string { s.scan(); return s.text() }
func (s *Scanner) Int() int       { return atoi(s.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }

func (w *Writer) Println(a ...interface{}) { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()                   { w.bf.Flush() }

type UnionFind struct {
	parent []int
	size   []int
}
