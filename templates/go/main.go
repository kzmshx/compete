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
	r, w := NewScanner(os.Stdin, maxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

// signed is a constraint that permits any signed integer type.
// int8, int16 are not included because they are rarely used.
type signed interface{ ~int | ~int32 | ~int64 }

// unsigned is a constraint that permits any unsigned integer type.
// uint8, uint16, uintptr are not included because they are rarely used.
type unsigned interface{ ~uint | ~uint32 | ~uint64 }

// integer is a constraint that permits any integer type.
type integer interface{ signed | unsigned }

// float is a constraint that permits any floating-point type.
type float interface{ ~float32 | ~float64 }

// actual is a constraint that permits any complex numeric type.
type actual interface{ integer | float }

// imaginary is a constraint that permits any complex numeric type.
type imaginary interface{ ~complex64 | ~complex128 }

// ordered is a constraint that permits any ordered type: any type
type ordered interface{ integer | float | ~string }

// addable is a constraint that permits any ordered type: any type
type addable interface {
	integer | float | imaginary | ~string
}

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
func Bin[T integer](n T) string { return strconv.FormatInt(int64(n), 2) }

// Oct returns the octal representation of n.
func Oct[T integer](n T) string { return strconv.FormatInt(int64(n), 8) }

// Hex returns the hexadecimal representation of n.
func Hex[T integer](n T) string { return strconv.FormatInt(int64(n), 16) }

// ParseInt converts s to int in base b.
func ParseInt(s string, b int) int { return int(Unwrap(strconv.ParseInt(s, b, 64))) }

// Digits returns the digits of n with base
func Digits[T integer](n T, base T) (r []T) {
	for n > 0 {
		r = append(r, n%base)
		n /= base
	}
	return r
}

// Max returns the maximum of a and b.
func Max[T ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of a and b.
func Min[T ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// ChooseMax sets the maximum value of a and b to a and returns the maximum value.
func ChooseMax[T ordered](a *T, b T) T {
	if *a < b {
		*a = b
	}
	return *a
}

// ChooseMin sets the minimum value of a and b to a and returns the minimum value.
func ChooseMin[T ordered](a *T, b T) T {
	if *a > b {
		*a = b
	}
	return *a
}

// Abs returns the absolute value of x.
func Abs[T actual](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

// Pow returns x**n, the base-x exponential of n.
func Pow[T actual](x T, n int) T {
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

func (s *Scanner) Int() int { return Atoi(s.String()) }

func (s *Scanner) Float64() float64 { return Atof(s.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }

func (w *Writer) Print(a ...interface{}) { fmt.Fprint(w.bf, a...) }

func (w *Writer) Printf(format string, a ...interface{}) { fmt.Fprintf(w.bf, format, a...) }

func (w *Writer) Println(a ...interface{}) { fmt.Fprintln(w.bf, a...) }

func (w *Writer) Flush() { w.bf.Flush() }

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
