package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func Solve(r *Scanner, w *Writer) {
	date := r.String()
	t := unwrap(time.Parse("2006/01/02", date))
	for t.Year()%(int(t.Month())*t.Day()) != 0 {
		t = t.AddDate(0, 0, 1)
	}
	w.Println(t.Format("2006/01/02"))
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

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }

func (w *Writer) Println(a ...interface{}) { fmt.Fprintln(w.bf, a...) }

func (w *Writer) Flush() { w.bf.Flush() }

type UnionFind struct {
	parent []int
	size   []int
}
