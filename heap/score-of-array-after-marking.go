package main

// You are given an array nums consisting of positive integers.
// Starting with score = 0, apply the following algorithm:
// Choose the smallest integer of the array that is not marked. If there is a tie, choose the one with the smallest index.
// Add the value of the chosen integer to score.
// Mark the chosen element and its two adjacent elements if they exist.
// Repeat until all the array elements are marked.
// Return the score you get after applying the above algorithm.

func findScore(nums []int) int64 {
    n := len(nums) 
    var score int64 = 0 
    marked := make([]int,n+2)
    pq := &MinHeap{}
	heap.Init(pq)
    for i := 0 ; i < n ; i++ {
        heap.Push(pq,pair{nums[i],i+1})
    }
    for pq.Len() > 0 {
        top := heap.Pop(pq).(pair)
        if marked[top.index] == 0 { 
            score += int64(top.value)
            marked[top.index] = 1
            marked[top.index+1] = 1
            marked[top.index-1] = 1
        }
    }
    return score
}

type pair struct {
    value int
    index int
}

type MinHeap []pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { 
    if h[i].value < h[j].value { return true }
    if h[i].value == h[j].value { return h[i].index < h[j].index }
    return false
}
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}
