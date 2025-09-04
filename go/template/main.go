package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Reader, w *Writer) {
	w.Println(r.Int())
}

func main() {
	r, w := NewReader(os.Stdin, 1*1024*1024), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

// ================================================================
// Constraints
// ================================================================

type Sgn interface{ ~int | ~int32 | ~int64 }
type Uns interface{ ~uint | ~uint32 | ~uint64 }
type Int interface{ Sgn | Uns }
type Float interface{ ~float32 | ~float64 }
type Num interface{ Int | Float }
type Ord interface{ Int | Float | ~string }

// ================================================================
// Conversion
// ================================================================

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
func Atoi(s string) int            { return Unwrap(strconv.Atoi(s)) }
func Atof(s string) float64        { return Unwrap(strconv.ParseFloat(s, 64)) }
func Itoa(i int) string            { return strconv.Itoa(i) }
func Bin[T Int](n T) string        { return strconv.FormatInt(int64(n), 2) }
func Oct[T Int](n T) string        { return strconv.FormatInt(int64(n), 8) }
func Hex[T Int](n T) string        { return strconv.FormatInt(int64(n), 16) }
func ParseInt(s string, b int) int { return int(Unwrap(strconv.ParseInt(s, b, 64))) }

// ================================================================
// IO
// ================================================================

type Reader struct{ sc *bufio.Scanner }

func NewReader(r io.Reader, size int) *Reader {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Reader{sc}
}
func (r *Reader) Bytes() []byte { r.sc.Scan(); return r.sc.Bytes() }
func (r *Reader) Int() int      { return Atoi(r.String()) }
func (r *Reader) Ints(n int) []int {
	return MakeSlice(n, func(i int) int { return r.Int() })
}
func (r *Reader) String() string { r.sc.Scan(); return r.sc.Text() }
func (r *Reader) Strings(n int) []string {
	return MakeSlice(n, func(i int) string { return r.String() })
}
func (r *Reader) Float64() float64 { return Atof(r.String()) }
func (r *Reader) Float64s(n int) []float64 {
	return MakeSlice(n, func(i int) float64 { return r.Float64() })
}

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer                  { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Print(a ...any)                     { fmt.Fprint(w.bf, a...) }
func (w *Writer) Printf(format string, a ...any)     { fmt.Fprintf(w.bf, format, a...) }
func (w *Writer) Println(a ...any)                   { fmt.Fprintln(w.bf, a...) }
func (w *Writer) PrintlnFloat64(a float64, prec int) { w.Printf("%.*f", prec, a) }
func (w *Writer) PrintlnYes(a bool) {
	if a {
		w.Println("Yes")
	} else {
		w.Println("No")
	}
}
func (w *Writer) Flush() { w.bf.Flush() }

// ================================================================
// Math
// ================================================================

// Digits は base 進数で n を表す数字の配列を返す
func Digits[T Int](n T, base T) (r []T) {
	for n > 0 {
		r = append(r, n%base)
		n /= base
	}
	return r
}

// Max は a と b の最大値を返す
func Max[T Ord](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min は a と b の最小値を返す
func Min[T Ord](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// ChooseMax は a と b の最大値を a に設定して返す
func ChooseMax[T Ord](a *T, b T) T {
	if *a < b {
		*a = b
	}
	return *a
}

// ChooseMin は a と b の最小値を a に設定して返す
func ChooseMin[T Ord](a *T, b T) T {
	if *a > b {
		*a = b
	}
	return *a
}

// Abs は x の絶対値を返す
func Abs[T Num](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

// Pow は x の n 乗を返す
func Pow[T Num](x T, n int) T {
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

// GCD は a と b の最大公約数を返す
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM は a と b の最小公倍数を返す
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// ================================================================
// Slices
// ================================================================

// MakeSlice は長さ n のスライスを作成し、各要素を f で初期化して返す
func MakeSlice[T any](n int, f func(i int) T) []T {
	a := make([]T, n)
	for i := range a {
		a[i] = f(i)
	}
	return a
}

// All は s のすべての要素が f を満たすかどうかを返す
func All[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

// Any は s のいずれかの要素が f を満たすかどうかを返す
func Any[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// Count は a の要素で f を満たすものの数を返す
func Count[T any](a []T, f func(T) bool) (count int) {
	for _, v := range a {
		if f(v) {
			count++
		}
	}
	return
}

// ================================================================
// Binary Search
// ================================================================

// BinarySearch は [l, r) の範囲で f が真となる最小の T を返す、f が真となる要素が存在しない場合は r を返す
func BinarySearch[T Int](l, r T, f func(T) bool) T {
	for l < r {
		m := T(uint(l+r) >> 1)
		if f(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// LowerBound は a[i] >= x となる最初の i を返す、そのような i が存在しない場合は n を返す
func LowerBound[T Ord](s []T, x T) int {
	return BinarySearch(0, len(s), func(i int) bool { return s[i] >= x })
}

// UpperBound は a[i] > x となる最初の i を返す、そのような i が存在しない場合は n を返す
func UpperBound[T Ord](s []T, x T) int {
	return BinarySearch(0, len(s), func(i int) bool { return s[i] > x })
}

// ================================================================
// CyclicInt 循環整数
// ================================================================

type CyclicInt[T Sgn] struct {
	v T
	c T
}

func NewCyclicInt[T Sgn](v, c T) CyclicInt[T] {
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

func (z CyclicInt[T]) Increment() CyclicInt[T] {
	return z.Advance(1)
}

func (z CyclicInt[T]) Decrement() CyclicInt[T] {
	return z.Advance(-1)
}

// ================================================================
// ModInt モジュラー整数
// ================================================================

type ModInt[T Int] struct {
	v T
	m T
}

func NewModInt[T Int](v, m T) ModInt[T] {
	v %= m
	if v < 0 {
		v += m
	}
	return ModInt[T]{v: v, m: m}
}

func (z ModInt[T]) assertModulus(x ModInt[T], op string) {
	if z.m != x.m {
		panic(fmt.Sprintf("ModInt: "+op+": modulus mismatch %v != %v", z.m, x.m))
	}
}

func (z ModInt[T]) Value() T {
	return z.v
}

func (z ModInt[T]) Add(x ModInt[T]) ModInt[T] {
	z.assertModulus(x, "Add")
	res := z.v + x.v
	if res >= z.m {
		res -= z.m
	}
	return ModInt[T]{v: res, m: z.m}
}

func (z ModInt[T]) Sub(x ModInt[T]) ModInt[T] {
	z.assertModulus(x, "Sub")
	res := z.v - x.v
	if res < 0 {
		res += z.m
	}
	return ModInt[T]{v: res, m: z.m}
}

func (z ModInt[T]) Mul(x ModInt[T]) ModInt[T] {
	z.assertModulus(x, "Mul")
	return ModInt[T]{v: (z.v * x.v) % z.m, m: z.m}
}

func (z ModInt[T]) Pow(n T) ModInt[T] {
	if n < 0 {
		return z.Inv().Pow(-n)
	}
	ret := NewModInt(1, z.m)
	base := z
	for n > 0 {
		if n&1 == 1 {
			ret = ret.Mul(base)
		}
		base = base.Mul(base)
		n >>= 1
	}
	return ret
}

func (z ModInt[T]) Inv() ModInt[T] {
	z.assertModulus(z, "Inv")
	return z.Pow(z.m - 2)
}

func (z ModInt[T]) Div(x ModInt[T]) ModInt[T] {
	z.assertModulus(x, "Div")
	return z.Mul(x.Inv())
}

func (z ModInt[T]) Equals(x ModInt[T]) bool {
	return z.v == x.v && z.m == x.m
}

// ================================================================
// Union Find Tree
// ================================================================

// UnionFindTree is a disjoint-set data structure.
type UnionFindTree struct {
	parent []int // parent[i] = parent of i
	size   []int // size[i] = number of elements in subtree rooted at i
}

// NewUnionFindTree creates a new union-find data structure with n elements.
func NewUnionFindTree(n int) *UnionFindTree {
	u := &UnionFindTree{parent: make([]int, n), size: make([]int, n)}
	for i := 0; i < n; i++ {
		u.parent[i], u.size[i] = -1, 1
	}
	return u
}

// Unite merges the components that elements x and y belong to.
func (u *UnionFindTree) Unite(x, y int) bool {
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

// Root returns the root of the component that element x belongs to.
func (u *UnionFindTree) Root(x int) int {
	// x is the root of the tree
	if u.parent[x] == -1 {
		return x
	}
	// Use path compression heuristic.
	u.parent[x] = u.Root(u.parent[x])
	return u.parent[x]
}

// Same returns true if elements x and y belong to the same component.
func (u *UnionFindTree) Same(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

// Size returns the size of the component that element x belongs to.
func (u *UnionFindTree) Size(x int) int {
	return u.size[u.Root(x)]
}

// ================================================================
// Priority Queue
// ================================================================

func Maximum[T Ord](lhs, rhs T) bool { return lhs < rhs }
func Minimum[T Ord](lhs, rhs T) bool { return lhs > rhs }

type priorityQueueItem[T any, P Ord] struct {
	value    T
	priority P
}

func newPriorityQueueItem[T any, P Ord](value T, priority P) *priorityQueueItem[T, P] {
	return &priorityQueueItem[T, P]{value: value, priority: priority}
}

type PriorityQueue[T any, P Ord] struct {
	items      []*priorityQueueItem[T, P]
	itemCount  uint
	comparator func(lhs, rhs P) bool
}

func NewPriorityQueue[T any, P Ord](heuristic func(lhs, rhs P) bool) *PriorityQueue[T, P] {
	items := make([]*priorityQueueItem[T, P], 1)
	items[0] = nil
	return &PriorityQueue[T, P]{items: items, itemCount: 0, comparator: heuristic}
}

func NewMaxPriorityQueue[T any, P Ord]() *PriorityQueue[T, P] {
	return NewPriorityQueue[T](Maximum[P])
}

func NewMinPriorityQueue[T any, P Ord]() *PriorityQueue[T, P] {
	return NewPriorityQueue[T](Minimum[P])
}

func (pq *PriorityQueue[T, P]) Size() uint {
	return pq.itemCount
}

func (pq *PriorityQueue[T, P]) Empty() bool {
	return pq.Size() == 0
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
	if pq.Empty() {
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

func (pq *PriorityQueue[T, P]) Peek() (value T, priority P, ok bool) {
	if pq.Empty() {
		ok = false
		return
	}
	return pq.items[1].value, pq.items[1].priority, true
}

// ================================================================
// Difference Array
// ================================================================

type Diff[T Num] struct{ delta []T }

func NewDiff[T Num](size int) *Diff[T] {
	return &Diff[T]{delta: make([]T, size+1)}
}

func (d *Diff[T]) Add(l, r int, val T) {
	d.delta[l] += val
	d.delta[r] -= val
}

func (d *Diff[T]) Increment(l, r int) {
	d.Add(l, r, 1)
}

func (d *Diff[T]) Build() []T {
	size := len(d.delta) - 1
	result := make([]T, size)
	var sum T
	for i := 0; i < size; i++ {
		sum += d.delta[i]
		result[i] = sum
	}
	return result
}

// ================================================================
// Vec 2次元ベクトル
// ================================================================

type Vec[T Sgn] struct{ X, Y T }

func NewVec[T Sgn](x, y T) Vec[T] { return Vec[T]{X: x, Y: y} }

func DirectionMap[K comparable, T Sgn](u, d, l, r K) map[K]Vec[T] {
	return map[K]Vec[T]{
		u: NewVec(T(0), T(-1)),
		d: NewVec(T(0), T(1)),
		l: NewVec(T(-1), T(0)),
		r: NewVec(T(1), T(0)),
	}
}

// Basic operations
func (v Vec[T]) Add(u Vec[T]) Vec[T]  { return NewVec(v.X+u.X, v.Y+u.Y) }
func (v Vec[T]) Sub(u Vec[T]) Vec[T]  { return NewVec(v.X-u.X, v.Y-u.Y) }
func (v Vec[T]) Scale(k T) Vec[T]     { return NewVec(v.X*k, v.Y*k) }
func (v Vec[T]) Equals(u Vec[T]) bool { return v.X == u.X && v.Y == u.Y }
func (v Vec[T]) InBounds(min, max Vec[T]) bool {
	return min.X <= v.X && v.X <= max.X && min.Y <= v.Y && v.Y <= max.Y
}
func (v Vec[T]) InGrid(h, w T) bool { return v.InBounds(NewVec(T(0), T(0)), NewVec(T(h), T(w))) }

// Distance
func (v Vec[T]) Chebyshev(u Vec[T]) T        { return Max(Abs(v.X-u.X), Abs(v.Y-u.Y)) }
func (v Vec[T]) EuclideanSquared(u Vec[T]) T { return Pow(v.X-u.X, 2) + Pow(v.Y-u.Y, 2) }
func (v Vec[T]) Manhattan(u Vec[T]) T        { return Abs(v.X-u.X) + Abs(v.Y-u.Y) }

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

func Intersect1D[T Ord](a, b [2]T) ([2]T, bool) {
	min, max := Max(a[0], b[0]), Min(a[1], b[1])
	return [2]T{min, max}, min <= max
}

func Intersect2D[T Ord](a, b [2][2]T) ([2][2]T, bool) {
	rowRange, okRowRange := Intersect1D(a[0], b[0])
	colRange, okColRange := Intersect1D(a[1], b[1])
	return [2][2]T{rowRange, colRange}, okRowRange && okColRange
}
