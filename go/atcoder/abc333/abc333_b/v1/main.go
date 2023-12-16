package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solve(r *Scanner, w *Writer) {
	S, T := r.String(), r.String()
	S1, S2 := int(S[0]), int(S[1])
	if S1 > S2 {
		S1, S2 = S2, S1
	}
	T1, T2 := int(T[0]), int(T[1])
	if T1 > T2 {
		T1, T2 = T2, T1
	}
	if ((S2-S1 == 1 || S2-S1 == 4) && (T2-T1 == 1 || T2-T1 == 4)) || ((S2-S1 == 3 || S2-S1 == 2) && (T2-T1 == 3 || T2-T1 == 2)) {
		w.Println("Yes")
	} else {
		w.Println("No")
	}
}

func main() {
	r, w := NewScanner(os.Stdin, maxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
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
