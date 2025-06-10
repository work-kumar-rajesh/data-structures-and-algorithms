func postorderTraversal(root *TreeNode) []int {
    ans := make([]int,0)
    var helper func(*TreeNode)
    helper = func(curr *TreeNode) {
        if curr == nil { return }
        helper(curr.Left)
        helper(curr.Right)
        ans = append(ans,curr.Val)
    }
    helper(root)
    return ans
}