package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	N := r.Int()
	edges := make([][]int, N)
	for i := 0; i < N-1; i++ {
		u, v := r.Int()-1, r.Int()-1
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}

	degrees := make([]int, N)
	for i, neighbours := range edges {
		degrees[i] = len(neighbours)
	}

	todo := make([]int, 0, N)
	for i, neighbours := range edges {
		if len(neighbours) == 1 {
			if i == 0 {
				w.Println(1)
				return
			}
			todo = append(todo, i)
		}
	}

	counts := make([]int, N)
	for i := range counts {
		counts[i] = 1
	}

	visited := make([]bool, N)
	for len(todo) > 0 {
		node := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		visited[node] = true
		if node == 0 {
			continue
		}
		for _, neighbour := range edges[node] {
			if visited[neighbour] {
				continue
			}
			degrees[neighbour]--
			if degrees[neighbour] == 1 {
				todo = append(todo, neighbour)
			}
			counts[neighbour] += counts[node]
		}
	}

	countsOfRootNeighbours := make([]int, len(edges[0]))
	for i, neighbours := range edges[0] {
		countsOfRootNeighbours[i] = counts[neighbours]
	}
	sort.Ints(countsOfRootNeighbours)

	ans := 0
	for i := 0; i < len(countsOfRootNeighbours)-1; i++ {
		ans += countsOfRootNeighbours[i]
	}
	w.Println(ans + 1)
}

func main() {
	r, w := NewScanner(os.Stdin, maxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

func unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func atoi(s string) int { return unwrap(strconv.Atoi(s)) }

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
