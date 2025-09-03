package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Reader, w *Writer) {
	H, W := r.Int(), r.Int()
	g := make([][]byte, H)
	var start, goal [2]int
	for i := 0; i < H; i++ {
		g[i] = r.Bytes()
		for j := 0; j < W; j++ {
			switch g[i][j] {
			case 'S':
				start = [2]int{i, j}
			case 'G':
				goal = [2]int{i, j}
			}
		}
	}

	visited := make([][][]bool, H)
	for i := range visited {
		visited[i] = make([][]bool, W)
		for j := range visited[i] {
			visited[i][j] = make([]bool, 2)
		}
	}

	type state struct {
		pos   [2]int
		door  int
		steps int
	}
	posEquals := func(a, b [2]int) bool {
		return a[0] == b[0] && a[1] == b[1]
	}
	posMove := func(a [2]int, d [2]int) [2]int {
		return [2]int{a[0] + d[0], a[1] + d[1]}
	}

	queue := []state{{start, 0, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if posEquals(cur.pos, goal) {
			w.Println(cur.steps)
			return
		}
		for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			pos := posMove(cur.pos, d)
			if pos[0] < 0 || pos[0] >= H || pos[1] < 0 || pos[1] >= W {
				continue
			}

			door := cur.door
			switch g[pos[0]][pos[1]] {
			case '#':
				continue
			case 'x':
				if cur.door == 0 {
					continue
				}
			case 'o':
				if cur.door == 1 {
					continue
				}
			case '?':
				door = 1 - door
			}

			if !visited[pos[0]][pos[1]][door] {
				visited[pos[0]][pos[1]][door] = true
				queue = append(queue, state{pos, door, cur.steps + 1})
			}
		}
	}

	w.Println(-1)
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
func (r *Reader) Bytes() []byte  { r.sc.Scan(); return r.sc.Bytes() }
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
