package heap

import (
	"testing"
)

func TestHeap_Push(t *testing.T) {
	heap := New[int]()
	heap.Push(5)
	heap.Push(3)
	heap.Push(7)

	// Pop() should return the largest value
	if heap.Top() != 7 {
		t.Error("Top() should return 7")
	}
	if heap.Pop() != 7 {
		t.Error("Pop() should return 7")
	}

	// Insert new value
	heap.Push(35)
	if heap.Pop() != 35 {
		t.Error("Pop() should return 35")
	}
	if heap.Pop() != 5 {
		t.Error("Pop() should return 5")
	}
	if heap.Pop() != 3 {
		t.Error("Pop() should return 3")
	}

	// When heap is empty
	defer func() {
		if r := recover(); r == nil {
			t.Error("Pop() should panic")
		}
	}()
	heap.Pop()
}

func BenchmarkHeap_1(b *testing.B) {
	heap := New[int]()
	for i := 0; i < b.N; i++ {
		heap.Push(i)
	}
}

func BenchmarkHeap_100(b *testing.B) {
	heap := New[int]()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			heap.Push(j)
		}
	}
}

func BenchmarkHeap_10000(b *testing.B) {
	heap := New[int]()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			heap.Push(j)
		}
	}
}

func BenchmarkHeap_1000000(b *testing.B) {
	heap := New[int]()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			heap.Push(j)
		}
	}
}
