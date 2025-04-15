package main

import (
	"container/heap"
	"math"
)

// PriorityQueue is a min-heap for Dijkstra's algorithm
type Item struct {
	node int
	dist int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Dijkstra's algorithm to find the shortest path and distance
func dijkstra(graph [][][2]int, start int, target int) ([]int, int) {
	n := len(graph)
	dist := make([]int, n)
	parent := make([]int, n)

	// Initialize distances to infinity and parents to -1
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
		parent[i] = -1
	}

	// Create a priority queue and insert the start node
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, dist: 0})
	dist[start] = 0

	// Dijkstra's main loop
	for pq.Len() > 0 {
		// Get the node with the smallest distance
		currItem := heap.Pop(pq).(*Item)
		currNode := currItem.node

		// If we have reached the target node, we can stop
		if currNode == target {
			break
		}

		// Relax the edges of the current node
		for _, edge := range graph[currNode] {
			neighbor := edge[0]
			weight := edge[1]

			if dist[currNode]+weight < dist[neighbor] {
				dist[neighbor] = dist[currNode] + weight
				parent[neighbor] = currNode
				heap.Push(pq, &Item{node: neighbor, dist: dist[neighbor]})
			}
		}
	}

	// If target node is unreachable, return nil path and -1 distance
	if dist[target] == math.MaxInt32 {
		return nil, -1
	}

	// Reconstruct the shortest path
	path := []int{}
	for node := target; node != -1; node = parent[node] {
		path = append([]int{node}, path...)
	}

	return path, dist[target]
}
