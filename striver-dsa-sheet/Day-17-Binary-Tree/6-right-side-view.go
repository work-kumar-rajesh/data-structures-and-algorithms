//Iterative 

func rightSideView(root *TreeNode) []int {
    ans := []int{}
    queue := make([]*TreeNode,0)
    if root !=nil {
        queue = append(queue,root)
    }
    for len(queue) > 0 {
        sz := len(queue)
        ans = append(ans,queue[sz-1].Val)
        for i := 0 ; i < sz ; i++ {
            top := queue[0]
            queue = queue[1:]
            if top.Left != nil { queue = append(queue,top.Left)}
            if top.Right != nil { queue = append(queue,top.Right)}
        }
    }
    return ans
}

//Recursive
func rightSideView(root *TreeNode) []int {
    var res []int

    var dfs func(node *TreeNode, depth int)
    dfs = func(node *TreeNode, depth int) {
        if node == nil {
            return
        }

        // If this is the first node we're visiting at this depth, add it
        if depth == len(res) {
            res = append(res, node.Val)
        }

        // Visit right subtree first to prioritize right-side nodes
        dfs(node.Right, depth+1)
        dfs(node.Left, depth+1)
    }

    dfs(root, 0)
    return res
}
