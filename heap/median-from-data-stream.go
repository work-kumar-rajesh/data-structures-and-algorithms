package main

import "container/heap"

type MedianFinder struct {
	h1 *MinHeap
	h2 *MaxHeap
}

func Constructor() MedianFinder {
	h1 := &MinHeap{}
	h2 := &MaxHeap{}
	heap.Init(h1)
	heap.Init(h2)
	return MedianFinder{h1, h2}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.h2, num)
	heap.Push(this.h1, heap.Pop(this.h2))
	if this.h2.Len() < this.h1.Len() {
		heap.Push(this.h2, heap.Pop(this.h1))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.h1.Len() == this.h2.Len() {
		return float64((*this.h1)[0]+(*this.h2)[0]) / 2.0
	}
	return float64((*this.h2)[0])
}

// implementations
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}
