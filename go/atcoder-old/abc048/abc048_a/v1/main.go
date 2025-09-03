package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	r := NewScanner(os.Stdin, MaxBufferSize)
	w := NewWriter(os.Stdout)
	defer w.Flush()

	a, s, c := r.String(), r.String(), r.String()
	w.Print(string(a[0]))
	w.Print(string(s[0]))
	w.Println(string(c[0]))
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

func (s *Scanner) Float64() float64 {
	n, err := strconv.ParseFloat(s.String(), 64)
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

// UnionFind is a disjoint-set data structure.
type UnionFind struct {
	parent []int // parent[i] = parent of i
	size   []int // size[i] = number of elements in subtree rooted at i
}

// NewUnionFind creates a new union-find data structure with n elements.
func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
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
func (u *UnionFind) Root(x int) int {
	// x is the root of the tree
	if u.parent[x] == -1 {
		return x
	}

	// Use path compression heuristic.
	u.parent[x] = u.Root(u.parent[x])
	return u.parent[x]
}

// IsSameSet returns true if elements x and y belong to the same component.
func (u *UnionFind) IsSameSet(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

// Union merges the components that elements x and y belong to.
func (u *UnionFind) Union(x, y int) bool {
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
func (u *UnionFind) Size(x int) int {
	return u.size[u.Root(x)]
}
