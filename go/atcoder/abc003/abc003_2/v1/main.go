package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Solve(r *Scanner, w *Writer) {
	S, T := r.String(), r.String()
	for i := 0; i < len(S); i++ {
		if S[i] == T[i] {
			continue
		}
		if S[i] == '@' && strings.Contains("atcoder", string(T[i])) {
			continue
		}
		if T[i] == '@' && strings.Contains("atcoder", string(S[i])) {
			continue
		}
		w.Println("You will lose")
		return
	}
	w.Println("You can win")
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
