package main

func kthSmallest(root *TreeNode, k int) int {
    result := -1
    bfs(root, &k, &result)
    return result
}

func bfs(root *TreeNode, counter, result *int) {
    if root == nil {
        return
    }

    bfs(root.Left, counter, result)

    *counter--
    if *counter == 0 {
        *result = root.Val
        return 
    }
    bfs(root.Right, counter, result)
}
