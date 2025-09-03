package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	pq := NewMinPriorityQueue[any, int]()
	Q := r.Int()
	for i := 0; i < Q; i++ {
		switch t := r.Int(); t {
		case 1:
			pq.Push(nil, r.Int())
		case 2:
			_, p, ok := pq.Pop()
			if !ok {
				panic("empty")
			}
			w.Println(p)
		default:
			panic("unexpected type " + Itoa(t))
		}
	}
}

func main() {
	r, w := NewScanner(os.Stdin, MaxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

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

func Itoa(i int) string { return strconv.Itoa(i) }

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

type UnionFind struct {
	parent []int
	size   []int
}

type priorityQueueItem[T any, P Ordered] struct {
	value    T
	priority P
}

func newPriorityQueueItem[T any, P Ordered](value T, priority P) *priorityQueueItem[T, P] {
	return &priorityQueueItem[T, P]{value: value, priority: priority}
}

func Minimum[T Ordered](lhs, rhs T) bool { return lhs > rhs }

type PriorityQueue[T any, P Ordered] struct {
	items      []*priorityQueueItem[T, P]
	itemCount  uint
	comparator func(lhs, rhs P) bool
}

func NewPriorityQueue[T any, P Ordered](heuristic func(lhs, rhs P) bool) *PriorityQueue[T, P] {
	items := make([]*priorityQueueItem[T, P], 1)
	items[0] = nil
	return &PriorityQueue[T, P]{items: items, itemCount: 0, comparator: heuristic}
}

func NewMinPriorityQueue[T any, P Ordered]() *PriorityQueue[T, P] {
	return NewPriorityQueue[T](Minimum[P])
}

func (pq *PriorityQueue[T, P]) Size() uint {
	return pq.itemCount
}

func (pq *PriorityQueue[T, P]) less(lhs, rhs uint) bool {
	return pq.comparator(pq.items[lhs].priority, pq.items[rhs].priority)
}

func (pq *PriorityQueue[T, P]) swap(lhs, rhs uint) {
	pq.items[lhs], pq.items[rhs] = pq.items[rhs], pq.items[lhs]
}

func (pq *PriorityQueue[T, P]) sink(k uint) {
	for 2*k <= pq.Size() {
		j := 2 * k
		if j < pq.Size() && pq.less(j, j+1) {
			j++
		}
		if !pq.less(k, j) {
			break
		}
		pq.swap(k, j)
		k = j
	}
}

func (pq *PriorityQueue[T, P]) swim(k uint) {
	for k > 1 && pq.less(k/2, k) {
		pq.swap(k/2, k)
		k /= 2
	}
}

func (pq *PriorityQueue[T, P]) Push(value T, priority P) {
	pq.items = append(pq.items, newPriorityQueueItem(value, priority))
	pq.itemCount++
	pq.swim(pq.Size())
}

func (pq *PriorityQueue[T, P]) Pop() (value T, priority P, ok bool) {
	if pq.Size() < 1 {
		ok = false
		return
	}
	max := pq.items[1]
	pq.swap(1, pq.Size())
	pq.items = pq.items[0:pq.Size()]
	pq.itemCount--
	pq.sink(1)
	return max.value, max.priority, true
}
