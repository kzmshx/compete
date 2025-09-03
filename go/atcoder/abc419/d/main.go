package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	_, m, s, t := r.Int(), r.Int(), r.String(), r.String()
	ops := make([][2]int, m)
	for i := range ops {
		ops[i] = [2]int{r.Int() - 1, r.Int() - 1}
	}
	diff := NewDiff[int](len(s))
	for _, op := range ops {
		diff.Increment(op[0], op[1]+1)
	}
	arr := diff.Build()
	ans := make([]byte, len(s))
	for i, v := range arr {
		if v%2 == 0 {
			ans[i] = s[i]
		} else {
			ans[i] = t[i]
		}
	}
	w.Println(string(ans))
}

func main() {
	r, w := NewScanner(os.Stdin, MaxBufferSize), NewWriter(os.Stdout)
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

func (s *Scanner) scan() bool { return s.sc.Scan() }

func (s *Scanner) text() string { return s.sc.Text() }

func (s *Scanner) String() string { s.scan(); return s.text() }

func (s *Scanner) Int() int { return Atoi(s.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }

func (w *Writer) Println(a ...any) { fmt.Fprintln(w.bf, a...) }

func (w *Writer) Flush() { w.bf.Flush() }

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

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

type UnionFind struct {
	parent []int
	size   []int
}

type priorityQueueItem[T any, P Ordered] struct {
	value    T
	priority P
}

type PriorityQueue[T any, P Ordered] struct {
	items      []*priorityQueueItem[T, P]
	itemCount  uint
	comparator func(lhs, rhs P) bool
}

type Diff[T Actual] struct {
	delta []T
}

func NewDiff[T Actual](size int) *Diff[T] {
	return &Diff[T]{delta: make([]T, size+1)}
}

func (d *Diff[T]) Add(l, r int, val T) {
	d.delta[l] += val
	d.delta[r] -= val
}

func (d *Diff[T]) Increment(l, r int) {
	d.Add(l, r, 1)
}

func (d *Diff[T]) Build() []T {
	size := len(d.delta) - 1
	result := make([]T, size)
	var sum T
	for i := 0; i < size; i++ {
		sum += d.delta[i]
		result[i] = sum
	}
	return result
}
