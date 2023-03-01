package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func dfs(v int, visited []bool, graph [][]int, component *[]int) {
	visited[v] = true
	*component = append(*component, v)
	for _, u := range graph[v] {
		if !visited[u] {
			dfs(u, visited, graph, component)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Считываем количество вершин и ребер
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	// Инициализируем граф
	graph := make([][]int, n)
	for i := 0; i < m; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		graph[a-1] = append(graph[a-1], b-1)
		graph[b-1] = append(graph[b-1], a-1)
	}

	// Запускаем DFS для поиска компоненты связности, содержащей первую вершину
	visited := make([]bool, n)
	component := []int{}
	dfs(0, visited, graph, &component)

	// Сортируем вершины компоненты связности
	sort.Ints(component)

	// Выводим результат
	fmt.Println(len(component))
	result := make([]string, len(component))
	for i, v := range component {
		result[i] = strconv.Itoa(v + 1)
	}
	fmt.Println(strings.Join(result, " "))
}
