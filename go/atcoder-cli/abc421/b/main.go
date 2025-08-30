package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func FibRev(n uint, memo map[uint]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	v := ParseInt(string(Reverse([]rune(Dec(FibRev(n-1, memo)+FibRev(n-2, memo))))), 10)
	memo[n] = v
	return v
}

func Solve(r *Scanner, w *Writer) {
	x, y := r.Int(), r.Int()
	memo := map[uint]int{}
	memo[1] = x
	memo[2] = y
	w.Println(FibRev(10, memo))
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
func (s *Scanner) scan() bool     { return s.sc.Scan() }
func (s *Scanner) text() string   { return s.sc.Text() }
func (s *Scanner) String() string { s.scan(); return s.text() }
func (s *Scanner) Int() int       { return Atoi(s.String()) }

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

func Atoi(s string) int            { return Unwrap(strconv.Atoi(s)) }
func Dec[T Integer](n T) string    { return strconv.FormatInt(int64(n), 10) }
func ParseInt(s string, b int) int { return int(Unwrap(strconv.ParseInt(s, b, 64))) }

func Reverse[T any](s []T) []T {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
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
