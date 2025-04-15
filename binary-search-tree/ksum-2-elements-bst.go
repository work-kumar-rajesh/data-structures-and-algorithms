package main

// Given the root of a binary search tree and an integer k, return true if there exist two elements
//  in the BST such that their sum is equal to k, or false otherwise.

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, k, map[int]struct{}{})
}

func dfs(root *TreeNode, k int, m map[int]struct{}) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return true
	}
	m[root.Val] = struct{}{}
	return dfs(root.Left, k, m) || dfs(root.Right, k, m)
}
