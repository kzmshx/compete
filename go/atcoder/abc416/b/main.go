package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solve(r *Reader, w *Writer) {
	s := []byte(r.String())
	ok := true
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			ok = true
		} else if ok {
			s[i] = 'o'
			ok = false
		}
	}
	w.Println(string(s))
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

type Reader struct{ sc *bufio.Scanner }

func NewReader(r io.Reader, size int) *Reader {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Reader{sc}
}
func (r *Reader) String() string { r.sc.Scan(); return r.sc.Text() }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Println(a ...any)  { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()            { w.bf.Flush() }
