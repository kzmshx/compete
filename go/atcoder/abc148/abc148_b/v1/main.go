package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	N, S, T := r.Int(), r.String(), r.String()
	for i := 0; i < N; i++ {
		w.Print(string(S[i]))
		w.Print(string(T[i]))
	}
	w.Println()
}

func main() {
	r := NewScanner(os.Stdin, MaxBufferSize)
	w := NewWriter(os.Stdout)
	defer w.Flush()

	Solve(r, w)
}

const MaxBufferSize = 1 * 1024 * 1024

type Scanner struct {
	sc *bufio.Scanner
}

func NewScanner(r io.Reader, size int) *Scanner {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}

func (s *Scanner) String() string {
	s.sc.Scan()
	return s.sc.Text()
}

func (s *Scanner) Int() int {
	n, err := strconv.Atoi(s.String())
	if err != nil {
		panic(err)
	}
	return n
}

type Writer struct {
	bf *bufio.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{bufio.NewWriter(w)}
}

func (w *Writer) Print(a ...interface{}) {
	fmt.Fprint(w.bf, a...)
}

func (w *Writer) Println(a ...interface{}) {
	fmt.Fprintln(w.bf, a...)
}

func (w *Writer) Flush() {
	w.bf.Flush()
}

type UnionFind struct {
	parent []int
	size   []int
}
