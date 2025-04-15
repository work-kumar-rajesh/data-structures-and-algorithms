package main

import (
	"fmt"
)

// Heap Operations and Their Internal Implementations
// The Go heap is implemented using a binary heap stored as an array. It provides the following core operations:

// Insertion (Push) → O(log n)
// Deletion (Pop) → O(log n)
// Heapify Up & Down (to maintain heap order property) → O(log n)
// The heap is a complete binary tree, where:
// The root (index 0) has the smallest element (for a min-heap).
// The parent-child relationship is maintained as:
// Parent at i → Left child at 2*i + 1, Right child at 2*i + 2.
// Child at i → Parent at (i-1)/2.

type MinHeapIMP struct {
	arr []int
}

func (h *MinHeapIMP) Push(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MinHeapIMP) Pop() int {
	if len(h.arr) == 0 {
		return -1 // Heap is empty
	}
	top := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1] // Move last element to root
	h.arr = h.arr[:len(h.arr)-1]   // Shrink array
	h.heapifyDown(0)
	return top
}

func (h *MinHeapIMP) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.arr[parent] > h.arr[index] {
			h.arr[parent], h.arr[index] = h.arr[index], h.arr[parent] // Swap
			index = parent
		} else {
			break
		}
	}
}

func (h *MinHeapIMP) heapifyDown(index int) {
	n := len(h.arr)
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < n && h.arr[left] < h.arr[smallest] {
			smallest = left
		}
		if right < n && h.arr[right] < h.arr[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.arr[index], h.arr[smallest] = h.arr[smallest], h.arr[index] // Swap
		index = smallest
	}
}

func main() {
	h := &MinHeapIMP{}
	h.Push(10)
	h.Push(5)
	h.Push(20)
	h.Push(2)
	fmt.Println(h.Pop()) // Output: 2
	fmt.Println(h.Pop()) // Output: 5
	fmt.Println(h.Pop()) // Output: 10
	fmt.Println(h.Pop()) // Output: 20
}
