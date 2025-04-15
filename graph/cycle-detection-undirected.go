package main

import "container/list"

func undirectedDfs(curr, par int, graph [][]int, vis []int) bool {
	vis[curr] = 1
	for _, node := range graph[curr] {
		if vis[node] == 0 {
			if undirectedDfs(node, curr, graph, vis) {
				return true
			}
		} else if node != par {
			return true
		}
	}
	return false
}

func undirectedCycleDetectionDFS(graph [][]int, n int) bool {
	vis := make([]int, n)
	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			if undirectedDfs(i, -1, graph, vis) {
				return true
			}
		}
	}
	return false
}

// ✅ Correct Time Complexity: O(N + E)
// ✅ Correct Space Complexity: O(N) (due to recursion stack + vis[] array).

//BFS

func isCycleBFS(graph [][]int, n int) bool {
	vis := make([]bool, n)

	for start := 0; start < n; start++ {
		if !vis[start] { // Only process unvisited components
			if bfsCheck(graph, start, vis) {
				return true // Cycle found
			}
		}
	}
	return false // No cycle found
}

func bfsCheck(graph [][]int, start int, vis []bool) bool {
	queue := list.New()
	queue.PushBack([2]int{start, -1}) // {node, parent}
	vis[start] = true

	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).([2]int)
		node, parent := front[0], front[1]

		for _, neighbor := range graph[node] {
			if !vis[neighbor] {
				vis[neighbor] = true
				queue.PushBack([2]int{neighbor, node})
			} else if neighbor != parent {
				return true // Cycle detected
			}
		}
	}
	return false
}

// Time Complexity:
// We visit each node once → O(V)
// We traverse all edges once → O(E)
// Total Complexity: O(V + E) ✅
// Space Complexity:
// Visited array: O(V)
// Queue storage (worst case all nodes in queue): O(V)
// Graph adjacency list: O(V + E)
// Total Complexity: O(V + E) ✅
