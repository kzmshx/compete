package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Query int

const (
	AddEdge Query = iota + 1
	FlipColor
	CheckReachability
)

func Solve(r *Reader, w *Writer) {
	n, q := r.Int(), r.Int()
	uf := NewUnionFindTree(n)
	rootBlackCount := make([]int, n)
	nodeIsBlack := make([]bool, n)

	for i := 0; i < q; i++ {
		switch Query(r.Int()) {
		case AddEdge:
			u, v := uf.Root(r.Int()-1), uf.Root(r.Int()-1)
			if u != v {
				if uf.Size(u) < uf.Size(v) {
					u, v = v, u
				}
				rootBlackCount[u] += rootBlackCount[v]
				uf.Unite(u, v)
			}
		case FlipColor:
			v := r.Int() - 1
			if nodeIsBlack[v] = !nodeIsBlack[v]; nodeIsBlack[v] {
				rootBlackCount[uf.Root(v)]++
			} else {
				rootBlackCount[uf.Root(v)]--
			}
		case CheckReachability:
			if v := uf.Root(r.Int() - 1); rootBlackCount[v] > 0 {
				w.Println("Yes")
			} else {
				w.Println("No")
			}
		default:
			panic("unexpected type")
		}
	}
}

func main() {
	r, w := NewReader(os.Stdin, 1*1024*1024), NewWriter(os.Stdout)
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

func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

type Reader struct{ sc *bufio.Scanner }

func NewReader(r io.Reader, size int) *Reader {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Reader{sc}
}
func (r *Reader) Int() int       { return Atoi(r.String()) }
func (r *Reader) String() string { r.sc.Scan(); return r.sc.Text() }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Println(a ...any)  { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()            { w.bf.Flush() }

type CyclicInt[T Signed] struct {
	v T
	c T
}

type ModInt[T Integer] struct {
	v T
	m T
}

type UnionFindTree struct {
	parent []int
	size   []int
}

func NewUnionFindTree(n int) *UnionFindTree {
	u := &UnionFindTree{parent: make([]int, n), size: make([]int, n)}
	for i := 0; i < n; i++ {
		u.parent[i], u.size[i] = -1, 1
	}
	return u
}

func (u *UnionFindTree) Unite(x, y int) bool {
	xRoot, yRoot := u.Root(x), u.Root(y)
	if xRoot == yRoot {
		return false
	}

	if u.size[xRoot] < u.size[yRoot] {
		xRoot, yRoot = yRoot, xRoot
	}
	u.parent[yRoot] = xRoot
	u.size[xRoot] += u.size[yRoot]
	return true
}

func (u *UnionFindTree) Root(x int) int {

	if u.parent[x] == -1 {
		return x
	}

	u.parent[x] = u.Root(u.parent[x])
	return u.parent[x]
}

func (u *UnionFindTree) Size(x int) int { return u.size[u.Root(x)] }

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
