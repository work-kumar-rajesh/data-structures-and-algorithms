func verticalTraversal(root *TreeNode) [][]int {
    type Pair struct {
        row int
        val int
    }

    pos := make(map[int][]Pair)
    var traverse func(node *TreeNode, row, col int)
    traverse = func(node *TreeNode, row, col int) {
        if node == nil {
            return
        }
        pos[col] = append(pos[col], Pair{row, node.Val})
        traverse(node.Left, row+1, col-1)
        traverse(node.Right, row+1, col+1)
    }

    traverse(root, 0, 0)


    cols := make([]int, 0, len(pos))
    for col := range pos {
        cols = append(cols, col)
    }
    sort.Ints(cols)

    var ans [][]int
    for _, col := range cols {
        nodes := pos[col]
        sort.Slice(nodes, func(i, j int) bool {
            if nodes[i].row != nodes[j].row {
                return nodes[i].row < nodes[j].row
            }
            return nodes[i].val < nodes[j].val
        })
        colVals := make([]int, len(nodes))
        for i, p := range nodes {
            colVals[i] = p.val
        }
        ans = append(ans, colVals)
    }

    return ans
}
