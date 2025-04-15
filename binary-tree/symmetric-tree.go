package main

func isSymmetricHelper(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isSymmetricHelper(root1.Left, root2.Right) && isSymmetricHelper(root1.Right, root2.Left)
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricHelper(root, root)
}
