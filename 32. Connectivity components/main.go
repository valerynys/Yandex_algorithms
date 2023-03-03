package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(node int, visited []bool, graph map[int][]int, component []int) {
	visited[node] = true
	component = append(component, node)
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(neighbor, visited, graph, component)
		}
	}
}

func findComponents(n int, edges [][]int) [][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	components := [][]int{}
	visited := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if !visited[i] {
			component := []int{}
			dfs(i, visited, graph, component)
			components = append(components, component)
		}
	}

	return components
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n := toInt(scanner.Text())
	scanner.Scan()
	m := toInt(scanner.Text())

	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		u := toInt(scanner.Text())
		scanner.Scan()
		v := toInt(scanner.Text())
		edges[i] = []int{u, v}
	}

	components := findComponents(n, edges)
	fmt.Println(len(components))
	for _, component := range components {
		fmt.Println(len(component))
		for _, vertex := range component {
			fmt.Print(vertex, " ")
		}
		fmt.Println()
	}/
}

func toInt(s string) int {
	res := 0
	for _, ch := range s {
		res = res*10 + int(ch-'0')
	}
	return res
}
