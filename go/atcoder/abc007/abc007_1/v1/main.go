package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	w.Println(r.Int() - 1)
}

func main() {
	r := NewScanner(os.Stdin, MaxBufferSize)
	w := NewWriter(os.Stdout)
	defer w.Flush()

	Solve(r, w)
}

const MaxBufferSize = 1 * 1024 * 1024

type Scanner struct{ sc *bufio.Scanner }

func NewScanner(r io.Reader, size int) *Scanner {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}
func (s *Scanner) scan() bool       { return s.sc.Scan() }
func (s *Scanner) text() string     { return s.sc.Text() }
func (s *Scanner) String() string   { s.scan(); return s.text() }
func (s *Scanner) Int() int         { return atoi(s.String()) }
func (s *Scanner) Float64() float64 { return atof(s.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer        { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Print(a ...interface{})   { fmt.Fprint(w.bf, a...) }
func (w *Writer) Println(a ...interface{}) { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()                   { w.bf.Flush() }

// atoi converts string to int.
func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

// atof converts string to float64.
func atof(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return n
}
