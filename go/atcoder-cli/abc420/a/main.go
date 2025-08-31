package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Reader, w *Writer) {
	x, y := NewCyclicInt(r.Int(), 12), r.Int()
	w.Println(x.Advance(y).Value())
}

func main() {
	r, w := NewReader(os.Stdin, MaxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

const MaxBufferSize = 1 * 1024 * 1024

type Reader struct{ sc *bufio.Scanner }

func NewReader(r io.Reader, size int) *Reader {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Reader{sc}
}
func (r *Reader) scan() bool     { return r.sc.Scan() }
func (r *Reader) text() string   { return r.sc.Text() }
func (r *Reader) String() string { r.scan(); return r.text() }
func (r *Reader) Int() int       { return Atoi(r.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Println(a ...any)  { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()            { w.bf.Flush() }

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

func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

type CyclicInt[T Signed] struct {
	v T
	c T
}

func NewCyclicInt[T Signed](v, c T) CyclicInt[T] {
	v = ((v-1)%c+c)%c + 1
	return CyclicInt[T]{v: v, c: c}
}

func (z CyclicInt[T]) Value() T {
	return z.v
}

func (z CyclicInt[T]) Advance(x T) CyclicInt[T] {
	v := ((z.v-1+x)%z.c+z.c)%z.c + 1
	return NewCyclicInt(v, z.c)
}

type ModInt[T Integer] struct {
	v T
	m T
}

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

type Diff[T Actual] struct{ delta []T }

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
