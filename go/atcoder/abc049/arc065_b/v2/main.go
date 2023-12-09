package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	r := NewScanner(os.Stdin)
	w := NewWriter(os.Stdout)
	defer w.Flush()

	N, K, L := r.Int(), r.Int(), r.Int()
	roadGroup := newUnionFind(N)
	railGroup := newUnionFind(N)
	for i := 0; i < K; i++ {
		p, q := r.Int()-1, r.Int()-1
		roadGroup.Union(p, q)
	}
	for i := 0; i < L; i++ {
		r, s := r.Int()-1, r.Int()-1
		railGroup.Union(r, s)
	}

	count := make(map[[2]int]int)
	for i := 0; i < N; i++ {
		count[[2]int{roadGroup.Root(i), railGroup.Root(i)}]++
	}

	for i := 0; i < N; i++ {
		if i > 0 {
			w.Print(" ")
		}
		w.Print(count[[2]int{roadGroup.Root(i), railGroup.Root(i)}])
	}
	w.Println()
}

const (
	BufferSize = 1 * 1024 * 1024
)

type Scanner struct {
	sc *bufio.Scanner
}

func NewScanner(r io.Reader) *Scanner {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, BufferSize), BufferSize)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}

func (s *Scanner) String() string {
	s.sc.Scan()
	return s.sc.Text()
}

func (s *Scanner) Int() int {
	n, _ := strconv.Atoi(s.String())
	return n
}

func (s *Scanner) Float64() float64 {
	n, _ := strconv.ParseFloat(s.String(), 64)
	return n
}

type Writer struct {
	buf *bufio.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{bufio.NewWriter(w)}
}

func (w *Writer) Print(a ...interface{}) {
	fmt.Fprint(w.buf, a...)
}

func (w *Writer) Println(a ...interface{}) {
	fmt.Fprintln(w.buf, a...)
}

func (w *Writer) Flush() {
	w.buf.Flush()
}

// unionFind is a disjoint-set data structure.
type unionFind struct {
	parent []int // parent[i] = parent of i
	size   []int // size[i] = number of elements in subtree rooted at i
}

// newUnionFind creates a new union-find data structure with n elements.
func newUnionFind(n int) *unionFind {
	u := &unionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.parent[i] = -1
		u.size[i] = 1
	}
	return u
}

// Root returns the root of the component that element x belongs to.
func (u *unionFind) Root(x int) int {
	// x is the root of the tree
	if u.parent[x] == -1 {
		return x
	}

	// Use path compression heuristic.
	u.parent[x] = u.Root(u.parent[x])
	return u.parent[x]
}

// IsSameSet returns true if elements x and y belong to the same component.
func (u *unionFind) IsSameSet(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

// Union merges the components that elements x and y belong to.
func (u *unionFind) Union(x, y int) bool {
	xRoot, yRoot := u.Root(x), u.Root(y)
	if xRoot == yRoot {
		return false
	}

	// Use union by size heuristic.
	// Merge smaller component into the larger one.
	if u.size[xRoot] < u.size[yRoot] {
		xRoot, yRoot = yRoot, xRoot
	}
	u.parent[yRoot] = xRoot
	u.size[xRoot] += u.size[yRoot]
	return true
}

// Size returns the size of the component that element x belongs to.
func (u *unionFind) Size(x int) int {
	return u.size[u.Root(x)]
}
