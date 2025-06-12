func findKthLargest(nums []int, k int) int {
    n := len(nums)
    freq := make(map[int]int)
    for i := 0; i < n ; i++ {
        freq[nums[i]]++
    }
    pq := &MaxHeap{}
    heap.Init(pq)
    for key,count := range freq {
        heap.Push(pq,[]int{key,count})
    }
    curr := 0 
    for pq.Len() > 0 {
        top := heap.Pop(pq).([]int)
        val , count := top[0] , top[1]
        curr = curr + count
        if curr >=k { return val }
    }
    return -1
}

type MaxHeap [][]int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}
