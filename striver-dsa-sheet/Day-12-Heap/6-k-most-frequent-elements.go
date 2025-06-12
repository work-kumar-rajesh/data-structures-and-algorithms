//USING HEAP O(N+klogN)
type freq struct {
    count int
    value int 
}

func topKFrequent(nums []int, k int) []int {
    n := len(nums) 
    counts := make(map[int]int)
    for i := 0 ; i < n ; i++ {
        counts[nums[i]]++
    }
    arr := make([]freq,0)
    for key, value := range counts {
        arr = append(arr,freq{value,key}) 
    }
    pq := &MaxHeap{}
    *pq = arr
    heap.Init(pq)
    ans := make([]int,0)
    for len(ans) < k {
        top := heap.Pop(pq).(freq)
        ans = append(ans,top.value)
    }
    return ans
}

type MaxHeap []freq 

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i,j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i,j int) { h[i],h[j] = h[j],h[i] }
func (h *MaxHeap) Push(x any) { *h = append(*h,x.(freq)) }
func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    val := old[n-1]
    *h = old[:n-1]
    return val
}


//EFFICENT (bucker sort ) (O(N))

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for _, val := range nums {
        m[val]++
    }
    bucket := make([][]int, len(nums)+1)
    for key, val := range m {
        bucket[val] = append(bucket[val], key)
    }
    ans := make([]int, 0, k)
    for i := len(bucket) - 1; i >=0; i-- {
        for _, val := range bucket[i] {
            if k > 0 {
                ans = append(ans, val)
                k--
            }
        }
    }
    return ans
}