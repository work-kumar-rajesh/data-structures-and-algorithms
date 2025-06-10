func topView(root *Node) []int {
    if root == nil { return nil }

    type pair struct {
        node *Node
        hd   int
    }

    queue := []pair{{root, 0}}
    hdMap := map[int]int{}

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:] // dequeue

		ok,_ := hdMap[curr.hd]
        if !ok { hdMap[curr.hd] = curr.node.Val }

        if curr.node.Left != nil {
            queue = append(queue, pair{curr.node.Left, curr.hd - 1})
        }
        if curr.node.Right != nil {
            queue = append(queue, pair{curr.node.Right, curr.hd + 1})
        }
    }

    keys := make([]int, 0, len(hdMap))
    for k := range hdMap {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    res := make([]int, 0, len(keys))
    for _, k := range keys {
        res = append(res, hdMap[k])
    }

    return res
}