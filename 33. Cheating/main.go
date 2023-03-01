package main

import (
	"fmt"
)

func dfs(node int, color []int, graph map[int][]int, c int) bool {
	color[node] = c
	for _, neighbor := range graph[node] {
		if color[neighbor] == 0 {
			if !dfs(neighbor, color, graph, -c) {
				return false
			}
		} else if color[neighbor] == c {
			return false
		}
	}
	return true
}

func divideStudents(n int, edges [][]int) string {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	color := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if color[i] == 0 {
			if !dfs(i, color, graph, 1) {
				return "NO"
			}
		}
	}
	return "YES"
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		edges[i] = []int{u, v}
	}
	fmt.Println(divideStudents(n, edges))
}
