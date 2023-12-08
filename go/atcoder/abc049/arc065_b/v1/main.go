package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	N, K, L := read[int](), read[int](), read[int]()
	roadGroup := newUnionFind(N)
	railGroup := newUnionFind(N)
	for i := 0; i < K; i++ {
		p, q := read[int]()-1, read[int]()-1
		roadGroup.Union(p, q)
	}
	for i := 0; i < L; i++ {
		r, s := read[int]()-1, read[int]()-1
		railGroup.Union(r, s)
	}

	count := make(map[[2]int]int)
	for i := 0; i < N; i++ {
		count[[2]int{roadGroup.Root(i), railGroup.Root(i)}]++
	}

	results := make([]int, N)
	for i := 0; i < N; i++ {
		root := [2]int{roadGroup.Root(i), railGroup.Root(i)}
		results[i] = count[root]
	}

	fmt.Println(sliceToString(results, " "))
}

// read reads a value from stdin.
func read[T any]() (r T) {
	fmt.Scan(&r)
	return r
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

func toString[T any](value T) string {
	switch v := any(value).(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	default:
		return fmt.Sprint(v)
	}
}

func sliceToString[T any](slice []T, sep string) string {
	var buffer bytes.Buffer
	for i, v := range slice {
		if i > 0 {
			buffer.WriteString(sep)
		}
		buffer.WriteString(toString(v))
	}
	return buffer.String()
}
