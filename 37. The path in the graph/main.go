package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Чтение входных данных
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	n := nextInt(scanner)

	// Создание графа
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = nextInt(scanner)
		}
	}

	start := nextInt(scanner) - 1
	end := nextInt(scanner) - 1

	// Выполнение алгоритма Дейкстры
	dist := make([]int, n)
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
		prev[i] = -1
	}

	dist[start] = 0
	queue := []int{start}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v := 0; v < n; v++ {
			if graph[u][v] == 1 {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					prev[v] = u
					queue = append(queue, v)
				}
			}
		}
	}

	// Вывод результата
	if dist[end] == -1 {
		fmt.Println(-1)
	} else {
		path := []int{end}
		for prev[path[0]] != -1 {
			path = append([]int{prev[path[0]]}, path...)
		}

		fmt.Println(dist[end])
		if dist[end] != 0 {
			for i := 0; i < len(path); i++ {
				fmt.Printf("%d ", path[i]+1)
			}
		}
	}
}

func nextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	var res int
	fmt.Sscanf(scanner.Text(), "%d", &res)
	return res
}
