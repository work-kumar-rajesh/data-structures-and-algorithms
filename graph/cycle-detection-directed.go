package main

import "container/list"

//DFS

// Instead of a separate dfsVis[] array, we can use vis[] to store 3 states:
// 0 → Unvisited
// 1 → Visiting (Recursion Stack)
// 2 → Fully Processed

func dfs(curr int, graph [][]int, vis []int) bool {
	vis[curr] = 1

	for _, node := range graph[curr] {
		if vis[node] == 0 {
			if dfs(node, graph, vis) {
				return true
			}
		} else if vis[node] == 1 {
			return true
		}
	}
	vis[curr] = 2 // Mark node as fully processed
	return false
}

func cycleDetectionDFS(graph [][]int, n int) bool {
	vis := make([]int, n)

	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			if dfs(i, graph, vis) {
				return true
			}
		}
	}
	return false
}

//BFS

func bfsCycleDetection(graph [][]int, n int) bool {
	indegree := make([]int, n)
	for i := 0; i < n; i++ {
		for _, neighbor := range graph[i] {
			indegree[neighbor]++
		}
	}
	queue := list.New()
	count := 0
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue.PushBack(i)
			count++
		}
	}
	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(int)
		for _, neighbor := range graph[curr] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue.PushBack(neighbor)
				count++
			}
		}
	}
	return count != n
}
