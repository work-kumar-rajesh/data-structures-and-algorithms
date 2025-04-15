package main

import "container/list"

//dfs

func bipartiteHelper(curr, markColor int, graph [][]int, color []int) bool {
	color[curr] = markColor

	for _, node := range graph[curr] {
		if color[node] == -1 {
			if !bipartiteHelper(node, 1-markColor, graph, color) {
				return false
			}
		} else if color[node] == markColor {
			return false
		}
	}
	return true
}

func bipartiteCheckDFS(graph [][]int, n int) bool {
	color := make([]int, n)
	for i := range color {
		color[i] = -1
	}
	for i := 0; i < n; i++ {
		if color[i] == -1 {
			if !bipartiteHelper(i, 0, graph, color) {
				return false
			}
		}
	}
	return true
}

//bfs

func bfsBipartiteHelper(idx int, graph [][]int, color []int) bool {
	queue := list.New()
	queue.PushBack([]int{idx, 0})
	for queue.Len() > 0 {
		front := queue.Remove(queue.Front())
		curr, col := front.([]int)[0], front.([]int)[1]
		for _, val := range graph[curr] {
			if color[val] == -1 {
				color[val] = 1 - col
				queue.PushBack([]int{val, color[val]})
			} else if color[val] == col {
				return false
			}
		}
	}
	return true
}

func bipartiteCheckBFS(graph [][]int, n int) bool {
	color := make([]int, n)
	for i := range color {
		color[i] = -1
	}
	for i := 0; i < n; i++ {
		if color[i] == -1 {
			if !bfsBipartiteHelper(i, graph, color) {
				return false
			}
		}
	}
	return true
}
