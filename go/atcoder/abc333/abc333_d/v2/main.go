package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	N := r.Int()
	var subtrees []int
	uf := NewUnionFind(N)
	for i := 0; i < N-1; i++ {
		u, v := r.Int()-1, r.Int()-1
		if u == 0 {
			subtrees = append(subtrees, v)
		} else {
			uf.Union(u, v)
		}
	}
	ans := math.MaxInt64
	for _, subtree := range subtrees {
		ans = min(ans, N-uf.Size(subtree))
	}
	w.Println(ans)
}

func main() {
	r, w := NewScanner(os.Stdin, maxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type integer interface {
	signed | unsigned
}

type float interface {
	~float32 | ~float64
}

type ordered interface {
	integer | float | ~string
}

func unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func atoi(s string) int { return unwrap(strconv.Atoi(s)) }

func min[T ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
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

func (s *Scanner) Int() int { return atoi(s.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }

func (w *Writer) Println(a ...interface{}) { fmt.Fprintln(w.bf, a...) }

func (w *Writer) Flush() { w.bf.Flush() }

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{parent: make([]int, n), size: make([]int, n)}
	for i := 0; i < n; i++ {
		u.parent[i], u.size[i] = -1, 1
	}
	return u
}

func (u *UnionFind) Union(x, y int) bool {
	xRoot, yRoot := u.Find(x), u.Find(y)
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

func (u *UnionFind) Find(x int) int {

	if u.parent[x] == -1 {
		return x
	}

	u.parent[x] = u.Find(u.parent[x])
	return u.parent[x]
}

func (u *UnionFind) Size(x int) int { return u.size[u.Find(x)] }
