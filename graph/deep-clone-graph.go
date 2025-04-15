package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)
	return cloneDFS(node, visited)
}

func cloneDFS(node *Node, visited map[int]*Node) *Node {
	if clonedNode, exists := visited[node.Val]; exists {
		return clonedNode
	}

	clonedNode := &Node{Val: node.Val}
	visited[node.Val] = clonedNode

	for _, neighbor := range node.Neighbors {
		clonedNode.Neighbors = append(clonedNode.Neighbors, cloneDFS(neighbor, visited))
	}

	return clonedNode
}

// ✅ O(N + E) → Each node and edge is processed exactly once.
