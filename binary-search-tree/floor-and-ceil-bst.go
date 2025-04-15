package main

func floor(root *TreeNode, key int) int {
	var res int = -1 // Assuming -1 indicates "no floor found"
	for root != nil {
		if root.Val == key {
			return root.Val
		} else if root.Val > key {
			root = root.Left
		} else {
			res = root.Val
			root = root.Right
		}
	}
	return res
}

func ceil(root *TreeNode, key int) int {
	var res int = -1 // Assuming -1 indicates "no ceil found"
	for root != nil {
		if root.Val == key {
			return root.Val
		} else if root.Val < key {
			root = root.Right
		} else {
			res = root.Val
			root = root.Left
		}
	}
	return res
}

// ✅ Time Complexity: O(h)
// ✅ Space Complexity: O(1)
