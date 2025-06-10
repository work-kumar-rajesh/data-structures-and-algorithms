func preorderTraversal(root *TreeNode) []int {
    ans := make([]int,0)
    var helper func(*TreeNode)
    helper = func(curr *TreeNode) {
        if curr == nil { return }
        ans = append(ans,curr.Val)
        helper(curr.Left)
        helper(curr.Right)
    }
    helper(root)
    return ans
}