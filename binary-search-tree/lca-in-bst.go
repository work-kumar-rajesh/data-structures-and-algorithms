package main

func lowestCommonAncestor0(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left // Both nodes are in the left subtree
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right // Both nodes are in the right subtree
		} else {
			return root // Split point, root is LCA
		}
	}
	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if (p.Val <= root.Val && q.Val >= root.Val) || (p.Val >= root.Val && q.Val <= root.Val) {
		return root
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor2(root.Left, p, q)
	} else {
		return lowestCommonAncestor2(root.Right, p, q)
	}

}
