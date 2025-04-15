package main

// recursive
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return 1 + max(left, right)
}

//iterative

type Pair struct {
	node   *TreeNode
	height int
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	stack := []Pair{Pair{root, 1}}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = max(ans, curr.height)
		if curr.node.Left != nil {
			stack = append(stack, Pair{curr.node.Left, curr.height + 1})
		}
		if curr.node.Right != nil {
			stack = append(stack, Pair{curr.node.Right, curr.height + 1})
		}
	}
	return ans
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

