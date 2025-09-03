package main

import (
	"fmt"
)

func main() {
	N, M := read[int](), read[int]()

	edges := make([][2]int, M)
	for i := 0; i < M; i++ {
		edges[i] = [2]int{read[int]() - 1, read[int]() - 1}
	}

	answer := 0

	// ex: index of the edge to exclude
	for ex := 0; ex < M; ex++ {
		// Union all edges except the one to exclude
		u := newUnionFind(N)
		for j := 0; j < M; j++ {
			if j == ex {
				continue
			}
			u.Union(edges[j][0], edges[j][1])
		}

		// Check if all nodes are connected
		for i := 1; i < N; i++ {
			if !u.IsSameSet(0, i) {
				answer++
				break
			}
		}
	}

	fmt.Println(answer)
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
