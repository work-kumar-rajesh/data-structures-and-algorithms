package main

func kthSmallest(root *TreeNode, k int) int {
	inOrder := make([]int, 0)
	inOrder = inOrderTraversal(root, inOrder)
	return inOrder[k-1]
}

func inOrderTraversal(root *TreeNode, inOrder []int) []int {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		inOrder = inOrderTraversal(root.Left, inOrder)
	}
	inOrder = append(inOrder, root.Val)
	if root.Right != nil {
		inOrder = inOrderTraversal(root.Right, inOrder)
	}
	return inOrder
}
