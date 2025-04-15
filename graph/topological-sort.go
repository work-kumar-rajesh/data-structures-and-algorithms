package main

import "container/list"

//DFS

func topoHelper(curr int, graph [][]int, vis, dfsVis []int, stack *[]int) bool {
	vis[curr] = 1
	dfsVis[curr] = 1 // Mark as part of current DFS path

	for _, neighbor := range graph[curr] {
		if vis[neighbor] == 0 { // Unvisited node
			if topoHelper(neighbor, graph, vis, dfsVis, stack) {
				return true // Cycle detected
			}
		} else if dfsVis[neighbor] == 1 { // Already in DFS path → Cycle found
			return true
		}
	}

	dfsVis[curr] = 0              // Remove from recursion stack
	*stack = append(*stack, curr) // Push to stack after processing
	return false
}

func topoSortDFS(graph [][]int, n int) []int {
	vis := make([]int, n)
	dfsVis := make([]int, n) // To track recursion path
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			if topoHelper(i, graph, vis, dfsVis, &stack) {
				return []int{} // Cycle detected, return empty list
			}
		}
	}

	// Reverse stack for topological order
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}

// BFS
func topoSortBfs(graph [][]int, n int) []int {
	ans := make([]int, 0)
	indegree := make([]int, n)
	for i := 0; i < n; i++ {
		for _, node := range graph[i] {
			indegree[node]++
		}
	}
	queue := list.New()
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue.PushBack(i)
		}
	}
	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(int)
		ans = append(ans, curr)
		for _, val := range graph[curr] {
			indegree[val]--
			if indegree[val] == 0 {
				queue.PushBack(val)
			}
		}
	}
	// If topological sort is incomplete, return an empty slice (cycle detected)
	if len(ans) != n {
		return []int{}
	}
	return ans
}
