package heap

import (
	"golang.org/x/exp/constraints"
)

type list[T constraints.Ordered] []T

func (l *list[T]) Len() int {
	return len(*l)
}

func (l *list[T]) Empty() bool {
	return l.Len() == 0
}

func (l *list[T]) Less(i, j int) bool {
	return (*l)[i] < (*l)[j]
}

func (l *list[T]) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *list[T]) Push(x T) {
	*l = append(*l, x)
}

func (l *list[T]) Pop() (x T) {
	x, *l = (*l)[len(*l)-1], (*l)[:len(*l)-1]
	return x
}

func (l *list[T]) Top() T {
	return (*l)[0]
}

type Heap[T constraints.Ordered] struct {
	heap list[T]
}

func New[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{}
}

func (h *Heap[T]) Push(value T) {
	// Push value to the end of the heap
	h.heap.Push(value)

	// From leaf to root
	i := h.heap.Len() - 1
	for i > 0 {
		// Heap structure ensures that parent index is (i - 1) / 2
		p := (i - 1) / 2

		// If parent is less than child, swap them
		if h.heap.Less(p, i) {
			h.heap.Swap(p, i)
		}

		i = p
	}
}

func (h *Heap[T]) Top() T {
	if h.heap.Empty() {
		panic("Heap is empty")
	}
	return h.heap.Top()
}

func (h *Heap[T]) Pop() T {
	if h.heap.Empty() {
		panic("Heap is empty")
	}

	// Pop the maximum value and move the minimum value to the top
	h.heap.Swap(0, h.heap.Len()-1)
	v := h.heap.Pop()

	// From root to leaf
	i := 0
	for i*2+1 < h.heap.Len() {
		// Heap structure ensures that child index is i*2 + 1 or i*2 + 2
		// Choose the larger child
		c := i*2 + 1
		if c+1 < h.heap.Len() && h.heap.Less(c, c+1) {
			c++
		}

		// Stop if parent is larger than child, otherwise swap parent and child
		if h.heap.Less(c, i) {
			break
		}
		h.heap.Swap(i, c)

		i = c
	}
	return v
}
