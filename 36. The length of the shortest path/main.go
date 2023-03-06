package main

import (
	"fmt"
)

func main() {

	// Read input
	var n, start, end int
	fmt.Scan(&n)

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&graph[i][j])
		}
	}

	fmt.Scan(&start, &end)

	// Find shortest path using BFS
	dist, prev := bfs(graph, start-1)

	if dist[end-1] == -1 {
		fmt.Println("-1")
	} else {
		// Reconstruct path
		path := make([]int, 0)
		at := end - 1
		for at != -1 {
			path = append(path, at+1)
			at = prev[at]
		}

		// Reverse path and print
		for i := len(path) - 1; i >= 0; i-- {
			fmt.Printf("%d ", path[i])
		}
		fmt.Println()
		fmt.Println(dist[end-1])
	}
}

func bfs(graph [][]int, start int) ([]int, []int) {
	n := len(graph)

	dist := make([]int, n)
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
		prev[i] = -1
	}

	dist[start] = 0

	queue := make([]int, 0)
	queue = append(queue, start)

	for len(queue) > 0 {
		at := queue[0]
		queue = queue[1:]

		for i := 0; i < n; i++ {
			if graph[at][i] == 1 && dist[i] == -1 {
				dist[i] = dist[at] + 1
				prev[i] = at
				queue = append(queue, i)
			}
		}
	}

	return dist, prev
}
