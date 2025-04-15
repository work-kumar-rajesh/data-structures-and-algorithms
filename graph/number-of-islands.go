package main

import "container/list"

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

//dfs

func islandHelper1(i, j, m, n int, grid [][]byte, vis [][]int) {
	if i < 0 || j < 0 || i == m || j == n || vis[i][j] == 1 || grid[i][j] == '0' {
		return
	}
	vis[i][j] = 1
	for k := 0; k < 4; k++ {
		islandHelper1(i+dx[k], j+dy[k], m, n, grid, vis)
	}
}

func numIslands1(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if vis[i][j] == 0 && grid[i][j] == '1' {
				ans++
				islandHelper1(i, j, m, n, grid, vis)
			}
		}
	}
	return ans
}

//bfs

func islandHelper(i, j, m, n int, grid [][]byte, vis [][]bool) {
	queue := list.New()
	queue.PushBack([]int{i, j})
	vis[i][j] = true // Mark as visited

	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).([]int)
		x, y := curr[0], curr[1]

		for k := 0; k < 4; k++ {
			nx, ny := x+dx[k], y+dy[k]
			if nx >= 0 && ny >= 0 && nx < m && ny < n && grid[nx][ny] == '1' && !vis[nx][ny] {
				queue.PushBack([]int{nx, ny})
				vis[nx][ny] = true // Mark as visited
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !vis[i][j] { // Found an unvisited island
				ans++
				islandHelper(i, j, m, n, grid, vis)
			}
		}
	}
	return ans
}
