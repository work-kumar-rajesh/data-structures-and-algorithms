package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	temp := root.Right
	root.Right = root.Left
	root.Left = nil
	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}
	curr.Right = temp
}
