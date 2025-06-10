func inorderTraversal(root *TreeNode) []int {
    ans := make([]int,0)
    var helper func(*TreeNode)
    helper = func(curr l*TreeNode) {
        if curr == nil { return }
        helper(curr.Left)
        ans = append(ans,curr.Val)
        helper(curr.Right)
    }
    helper(root)
    return ans
}