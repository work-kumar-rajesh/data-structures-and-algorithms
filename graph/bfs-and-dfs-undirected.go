package main

import "container/list"

// DFS
func dfsTraversal(node *Node, visited map[int]bool, result *[]int) {
	if node == nil || visited[node.Val] {
		return
	}

	visited[node.Val] = true
	*result = append(*result, node.Val) // Process the node

	for _, neighbor := range node.Neighbors {
		dfsTraversal(neighbor, visited, result)
	}
}

func dfsGraph(node *Node) []int {
	if node == nil {
		return []int{}
	}

	visited := make(map[int]bool)
	result := []int{}
	dfsTraversal(node, visited, &result)
	return result
}

// ✅ Time Complexity: O(N + E) (Visits each node and edge once)
// ✅ Space Complexity: O(N) (Recursive stack in worst case)

//BFS

func bfsGraph(node *Node) []int {
	if node == nil {
		return []int{}
	}

	visited := make(map[int]bool)
	queue := list.New()
	result := []int{}

	queue.PushBack(node)
	visited[node.Val] = true

	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(*Node)
		result = append(result, curr.Val) // Process the node

		for _, neighbor := range curr.Neighbors {
			if !visited[neighbor.Val] {
				queue.PushBack(neighbor)
				visited[neighbor.Val] = true
			}
		}
	}

	return result
}

// ✅ Time Complexity: O(N + E) (Processes each node and edge once)
// ✅ Space Complexity: O(N) (Queue holds at most all nodes)
