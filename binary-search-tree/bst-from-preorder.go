package main

import "math"

// Given an array of integers preorder, which represents the preorder traversal of a BST (i.e., binary search tree),
// construct the tree and return its root.
// It is guaranteed that there is always possible to find a binary search tree with the given requirements for the given test cases.
// A binary search tree is a binary tree where for every node, any descendant of Node.left has a value strictly less than
//  Node.val, and any descendant of Node.right has a value strictly greater than Node.val.
// A preorder traversal of a binary tree displays the value of the node first, then traverses Node.left, then traverses Node.right.

func bstFromPreorderHelper(start, end, n int, preorder []int) *TreeNode {
	if start > end {
		return nil
	}
	curr := &TreeNode{Val: preorder[start]}
	if start == end {
		return curr
	}
	index := end + 1
	for i := start + 1; i <= end; i++ {
		if preorder[i] > preorder[start] {
			index = i
			break
		}
	}
	curr.Left = bstFromPreorderHelper(start+1, index-1, n, preorder)
	curr.Right = bstFromPreorderHelper(index, end, n, preorder)
	return curr
}
func bstFromPreorder(preorder []int) *TreeNode {
	n := len(preorder)
	return bstFromPreorderHelper(0, n-1, n, preorder)
}

//optimized approach

func bstFromPreorder2(preorder []int) *TreeNode {
	index := 0
	return construct(preorder, &index, math.MinInt, math.MaxInt)
}

func construct(preorder []int, index *int, min, max int) *TreeNode {
	if *index >= len(preorder) {
		return nil
	}

	val := preorder[*index]
	if val < min || val > max {
		return nil
	}

	*index++ // Move to next element
	node := &TreeNode{Val: val}
	node.Left = construct(preorder, index, min, val)  // Left subtree
	node.Right = construct(preorder, index, val, max) // Right subtree
	return node
}
