package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsSearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	switch {
	case root.Val == val:
		return root
	case root.Val > val:
		return dfsSearchBST(root.Left, val)
	default:
		return dfsSearchBST(root.Right, val)
	}
}

func bfsSearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(*TreeNode)
		if curr.Val == val {
			return curr
		}
		if curr.Val > val && curr.Left != nil {
			queue.PushBack(curr.Left)
		}
		if curr.Val < val && curr.Right != nil {
			queue.PushBack(curr.Right)
		}
	}
	return nil
}
