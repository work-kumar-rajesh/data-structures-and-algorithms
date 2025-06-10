func widthOfBinaryTree(root *TreeNode) int {
    if root == nil { return 0 }
    type pair struct {
        *TreeNode
        hd int
    }
    queue := make([]pair,0)
    queue = append(queue,pair{root,1})
    ans := 0
    for len(queue) > 0 {
        sz := len(queue)
        ans = max(ans,queue[sz-1].hd - queue[0].hd+1)
        for i := 0 ; i < sz ; i++ {
            top := queue[0]
            queue = queue[1:]
            if top.Left != nil { queue = append(queue,pair{top.Left,2*top.hd-1})}
            if top.Right != nil { queue = append(queue,pair{top.Right,2*top.hd})}
        }
    }
    return ans
}