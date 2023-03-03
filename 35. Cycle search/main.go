package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	graph      [][]int
	visited    []int
	N          int
	flag       bool
	flagfind   bool
	res        []int
	startCycle int = -1
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	N, _ = strconv.Atoi(scan(scanner))
	graph = make([][]int, N+1)
	visited = make([]int, N+1)
	for i := 1; i <= N; i++ {
		graph[i] = make([]int, N+1)
		for j := 1; j <= N; j++ {
			graph[i][j], _ = strconv.Atoi(scan(scanner))
		}
	}

	for i := 1; i <= N; i++ {
		if visited[i] == 0 {
			dfs(i, -1)
		}
	}

	if len(res) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Println(len(res))
		for _, i := range res {
			fmt.Printf("%d ", i)
		}
	}
}

func dfs(current, prev int) {
	visited[current] = 1
	for i := 1; i <= N && !flagfind; i++ {
		if graph[current][i] == 1 && visited[i] == 0 {
			dfs(i, current)
			if flag && !flagfind {
				res = append(res, current)
				if current == startCycle {
					flagfind = true
				}
				return
			}
		} else if graph[current][i] == 1 && visited[i] == 1 && i != prev {
			startCycle = i
			res = append(res, current)
			flag = true
			return
		}
	}
}

func scan(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
