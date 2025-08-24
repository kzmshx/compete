package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	w.Println(r.Int())
}

func main() {
	r, w := NewScanner(os.Stdin, MaxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

// ================================================================
// IO
// ================================================================

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

func (s *Scanner) Float64() float64 { return Atof(s.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }

func (w *Writer) Print(a ...any) { fmt.Fprint(w.bf, a...) }

func (w *Writer) Printf(format string, a ...any) { fmt.Fprintf(w.bf, format, a...) }

func (w *Writer) Println(a ...any) { fmt.Fprintln(w.bf, a...) }

func (w *Writer) Flush() { w.bf.Flush() }

// ================================================================
// Constraints
// ================================================================

// Signed is a constraint that permits any Signed integer type.
// int8, int16 are not included because they are rarely used.
type Signed interface{ ~int | ~int32 | ~int64 }

// Unsigned is a constraint that permits any Unsigned integer type.
// uint8, uint16, uintptr are not included because they are rarely used.
type Unsigned interface{ ~uint | ~uint32 | ~uint64 }

// Integer is a constraint that permits any Integer type.
type Integer interface{ Signed | Unsigned }

// Float is a constraint that permits any floating-point type.
type Float interface{ ~float32 | ~float64 }

// Actual is a constraint that permits any complex numeric type.
type Actual interface{ Integer | Float }

// Imaginary is a constraint that permits any complex numeric type.
type Imaginary interface{ ~complex64 | ~complex128 }

// Ordered is a constraint that permits any Ordered type: any type
type Ordered interface{ Integer | Float | ~string }

// Addable is a constraint that permits any ordered type: any type
type Addable interface {
	Integer | Float | Imaginary | ~string
}

// ================================================================
// Conversion
// ================================================================

// Unwrap returns v if err is nil, otherwise panics with err.
func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// Atoi converts string s to int.
func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

// Atof converts string s to float64.
func Atof(s string) float64 { return Unwrap(strconv.ParseFloat(s, 64)) }

// Itoa converts int i to string.
func Itoa(i int) string { return strconv.Itoa(i) }

// Bin returns the binary representation of n.
func Bin[T Integer](n T) string { return strconv.FormatInt(int64(n), 2) }

// Oct returns the octal representation of n.
func Oct[T Integer](n T) string { return strconv.FormatInt(int64(n), 8) }

// Hex returns the hexadecimal representation of n.
func Hex[T Integer](n T) string { return strconv.FormatInt(int64(n), 16) }

// ParseInt converts s to int in base b.
func ParseInt(s string, b int) int { return int(Unwrap(strconv.ParseInt(s, b, 64))) }

// ================================================================
// Math
// ================================================================

// Digits returns the digits of n with base
func Digits[T Integer](n T, base T) (r []T) {
	for n > 0 {
		r = append(r, n%base)
		n /= base
	}
	return r
}

// Max returns the maximum of a and b.
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of a and b.
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// ChooseMax sets the maximum value of a and b to a and returns the maximum value.
func ChooseMax[T Ordered](a *T, b T) T {
	if *a < b {
		*a = b
	}
	return *a
}

// ChooseMin sets the minimum value of a and b to a and returns the minimum value.
func ChooseMin[T Ordered](a *T, b T) T {
	if *a > b {
		*a = b
	}
	return *a
}

// Abs returns the absolute value of x.
func Abs[T Actual](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

// Pow returns x**n, the base-x exponential of n.
func Pow[T Actual](x T, n int) T {
	y := T(1)
	for n > 0 {
		if n&1 == 1 {
			y *= x
		}
		x *= x
		n >>= 1
	}
	return y
}

// GCD returns the greatest common divisor of a and b.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of a and b.
func LCM(a, b int) int { return a * b / GCD(a, b) }

// ================================================================
// Slices
// ================================================================

// All checks if all elements in s satisfy f.
func All[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

// Any checks if any element in s satisfies f.
func Any[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// Contains checks if s contains e.
func Contains[T comparable](s []T, e T) bool {
	return Any(s, func(x T) bool { return x == e })
}

// LowerBound returns the first index i in [0, n) such that a[i] >= x.
// If there is no such index, it returns n.
func LowerBound[T Ordered](s []T, x T) int {
	l, r := 0, len(s)
	for l < r {
		m := int(uint(l+r) >> 1)
		if s[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

// UpperBound returns the first index i in [0, n) such that a[i] > x.
// If there is no such index, it returns n.
func UpperBound[T Ordered](s []T, x T) int {
	l, r := 0, len(s)
	for l < r {
		m := int(uint(l+r) >> 1)
		if s[m] <= x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

// ================================================================
// Union-Find
// ================================================================

// UnionFind is a disjoint-set data structure.
type UnionFind struct {
	parent []int // parent[i] = parent of i
	size   []int // size[i] = number of elements in subtree rooted at i
}

// NewUnionFind creates a new union-find data structure with n elements.
func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{parent: make([]int, n), size: make([]int, n)}
	for i := 0; i < n; i++ {
		u.parent[i], u.size[i] = -1, 1
	}
	return u
}

// Union merges the components that elements x and y belong to.
func (u *UnionFind) Union(x, y int) bool {
	xRoot, yRoot := u.Find(x), u.Find(y)
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

// Find returns the root of the component that element x belongs to.
func (u *UnionFind) Find(x int) int {
	// x is the root of the tree
	if u.parent[x] == -1 {
		return x
	}
	// Use path compression heuristic.
	u.parent[x] = u.Find(u.parent[x])
	return u.parent[x]
}

// IsSame returns true if elements x and y belong to the same component.
func (u *UnionFind) IsSame(x, y int) bool { return u.Find(x) == u.Find(y) }

// Size returns the size of the component that element x belongs to.
func (u *UnionFind) Size(x int) int { return u.size[u.Find(x)] }

// ================================================================
// Priority Queue
// ================================================================

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

// ================================================================
// Utilities
// ================================================================

// RandomString generates a random string of length n.
func RandomString(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}

// RenderGraph renders a graph in Mermaid format.
func RenderGraph(graph [][]int, root int) {
	filename := fmt.Sprintf("graph-%s.md", RandomString(8))
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	w := NewWriter(f)
	defer w.Flush()

	w.Println("```mermaid")
	w.Println("graph TD;")

	visited := make([]bool, len(graph))

	q := []int{root}
	visited[root] = true
	w.Printf("  %d((%d))\n", root, root)

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, n := range graph[v] {
			if visited[n] {
				continue
			}

			q = append(q, n)
			visited[n] = true
			w.Printf("  %d((%d))\n", n, n)
			w.Printf("  %d --- %d\n", v, n)
		}
	}

	w.Println("```")
}
