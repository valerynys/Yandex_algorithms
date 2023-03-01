package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	graph := make([][]int, n)
	inDegree := make([]int, n)

	for i := 0; i < m; i++ {
		scanner.Scan()
		parts := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		graph[a-1] = append(graph[a-1], b-1)
		inDegree[b-1]++
	}

	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0)

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u+1)

		for _, v := range graph[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(result) != n {
		fmt.Println("-1")
	} else {
		for _, v := range result {
			fmt.Printf("%d ", v)
		}
	}
}
