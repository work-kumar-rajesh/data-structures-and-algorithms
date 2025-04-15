package main

func maxPathSumHelper(node *TreeNode, ans *int) int {
	if node == nil {
		return -1e9
	}
	leftSum := maxPathSumHelper(node.Left, ans)
	rightSum := maxPathSumHelper(node.Right, ans)
	curr := node.Val + max(leftSum, rightSum)
	*ans = max(*ans, leftSum, rightSum, curr, node.Val, leftSum+rightSum+node.Val)
	return max(curr, node.Val)
}

func maxPathSum(root *TreeNode) int {
	var ans int = root.Val
	maxPathSumHelper(root, &ans)
	return ans
}
