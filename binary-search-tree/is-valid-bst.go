package main

func validator(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	switch {
	case root.Val <= min || root.Val >= max:
		return false
	default:
		return validator(root.Left, min, root.Val) && validator(root.Right, root.Val, max)
	}
}
func isValidBST(root *TreeNode) bool {
	return validator(root, -1e15, 1e15)
}
