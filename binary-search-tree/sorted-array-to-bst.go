package main

// Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
// A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

func bstHelper(start, end int, nums []int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = bstHelper(start, mid-1, nums)
	root.Right = bstHelper(mid+1, end, nums)
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	return bstHelper(0, len(nums)-1, nums)
}

// ⏳ Time Complexity: O(N)
// The function recursively processes each element exactly once, creating a new TreeNode for each number in nums.
// The recursive calls divide the array into two halves, but each element is processed only once.
// Total work done is proportional to the number of elements → O(N).
// Even though it looks like a divide-and-conquer approach (which often leads to O(N log N) complexity), each node is created once, making the overall complexity O(N).

// 🛠️ Space Complexity: O(N)
// 1. Recursive Call Stack: O(log N)
// Since the recursion depth is equal to the height of the BST, which is O(log N) for a balanced BST, the stack space required is O(log N).
// 2. Tree Nodes Storage: O(N)
// We create N nodes to store the BST, so the extra space for the tree itself is O(N).
// Total Space Complexity: O(N)
// O(log N) (call stack) + O(N) (tree nodes) = O(N).
// The dominant term is O(N), so we express the final complexity as O(N).
