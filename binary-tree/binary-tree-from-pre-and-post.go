package main

// Given two integer arrays, preorder and postorder where preorder is the preorder traversal of a binary
// tree of distinct values and postorder is the postorder traversal of the same tree, reconstruct and return the binary tree.
// If there exist multiple answers, you can return any of them.

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(i, j int, preorder, postorder []int, pos, pos2 map[int]int) *TreeNode {
	root := &TreeNode{
		Val:   preorder[i],
		Left:  nil,
		Right: nil,
	}
	if i == j {
		return root
	}
	x := pos[preorder[i]]
	y := pos2[postorder[x-1]]
	if i+1 <= y-1 {
		root.Left = helper(i+1, y-1, preorder, postorder, pos, pos2)
	}
	root.Right = helper(y, j, preorder, postorder, pos, pos2)
	return root
}
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)
	pos := make(map[int]int)
	pos2 := make(map[int]int)
	for i := 0; i < n; i++ {
		pos[postorder[i]] = i
	}
	for i := 0; i < n; i++ {
		pos2[preorder[i]] = i
	}
	return helper(0, n-1, preorder, postorder, pos, pos2)
}
