package main

import (
	"fmt"
)

func main() {
	// Get input
	N, M := read[int](), read[int]()
	edges := make([][2]int, M)
	for i := 0; i < M; i++ {
		edges[i] = [2]int{read[int]() - 1, read[int]() - 1}
	}

	disconnectedNodes := make([]int, M)

	// Union all edges from the end to the beginning
	uf := newUnionFind(N)
	for i := M - 1; i >= 0; i-- {
		a, b := edges[i][0], edges[i][1]
		if uf.IsSameSet(a, b) {
			// If a and b are already connected, the number of groups does not change
			disconnectedNodes[i] = 0
		} else {
			// If a and b are not connected, the number of groups increases by size(a) * size(b)
			disconnectedNodes[i] = uf.Size(a) * uf.Size(b)
		}
		uf.Union(a, b)
	}

	// Finally, cumulative sum of disconnected nodes is the answer
	sum := 0
	for _, v := range disconnectedNodes {
		sum += v
		fmt.Println(sum)
	}
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
