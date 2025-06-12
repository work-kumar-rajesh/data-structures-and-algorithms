type Item struct {
	value, arrayIndex, elementIndex int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value } 
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func mergeKSortedArrays(arrays [][]int) []int {
	result := []int{}

	h := &MinHeap{}

	for i := 0; i < len(arrays); i++ {
		if len(arrays[i]) > 0 {
			heap.Push(h, Item{value: arrays[i][0], arrayIndex: i, elementIndex: 0})
		}
	}

	for h.Len() > 0 {
		item := heap.Pop(h).(Item)

		result = append(result, item.value)
		if item.elementIndex+1 < len(arrays[item.arrayIndex]) {
			nextElement := arrays[item.arrayIndex][item.elementIndex+1]
			heap.Push(h, Item{
				value:        nextElement,
				arrayIndex:   item.arrayIndex,
				elementIndex: item.elementIndex + 1,
			})
		}
	}

	return result
}

