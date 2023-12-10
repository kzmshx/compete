package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	w.Println(r.Int())
}

func main() {
	r := NewScanner(os.Stdin, MaxBufferSize)
	w := NewWriter(os.Stdout)
	defer w.Flush()

	Solve(r, w)
}

// signed is a constraint that permits any signed integer type.
type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// unsigned is a constraint that permits any unsigned integer type.
type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// integer is a constraint that permits any integer type.
type integer interface {
	signed | unsigned
}

// float is a constraint that permits any floating-point type.
type float interface {
	~float32 | ~float64
}

// actual is a constraint that permits any complex numeric type.
type actual interface {
	integer | float
}

// imaginary is a constraint that permits any complex numeric type.
type imaginary interface {
	~complex64 | ~complex128
}

// ordered is a constraint that permits any ordered type: any type
type ordered interface {
	integer | float | ~string
}

// addable is a constraint that permits any ordered type: any type
type addable interface {
	integer | float | imaginary | ~string
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

// chmax sets the maximum value of a and b to a and returns the maximum value.
func chmax[T ordered](a *T, b T) T {
	if *a < b {
		*a = b
	}
	return *a
}

// chmin sets the minimum value of a and b to a and returns the minimum value.
func chmin[T ordered](a *T, b T) T {
	if *a > b {
		*a = b
	}
	return *a
}

// max returns the maximum value of a and b.
func max[T ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// min returns the minimum value of a and b.
func min[T ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// abs returns the absolute value of x.
func abs[T actual](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

// sliceFill returns a slice of length n with each element set to v.
func sliceFill[T any](n int, v T) []T {
	r := make([]T, n)
	for i := 0; i < n; i++ {
		r[i] = v
	}
	return r
}

// sliceFunc returns a slice of length n with each element set to the result of f(i).
func sliceFunc[T any](n int, f func(int) T) []T {
	r := make([]T, n)
	for i := 0; i < n; i++ {
		r[i] = f(i)
	}
	return r
}

// sliceReverse reverses the order of elements in s.
func sliceReverse[T any](s []T) []T {
	r := make([]T, len(s))
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = s[j], s[i]
	}
	return r
}

// sliceMax returns the maximum value of s.
func sliceMax[T ordered](s []T) (r T) {
	for _, v := range s {
		chmax(&r, v)
	}
	return r
}

// sliceSum returns the sum of s.
func sliceSum[T addable](s []T) (r T) {
	for _, v := range s {
		r += v
	}
	return r
}

// gcd returns the greatest common divisor of a and b.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm returns the least common multiple of a and b.
func lcm(a, b int) int {
	return a * b / gcd(a, b)
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
