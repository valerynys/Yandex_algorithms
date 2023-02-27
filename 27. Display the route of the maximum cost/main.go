package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Scan(&a[i][j])
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + a[i-1][j-1]
		}
	}

	i, j := n, m
	path := make([]byte, 0, n+m-2)
	for i > 1 || j > 1 {
		if dp[i-1][j] >= dp[i][j-1] {
			path = append(path, 'D')
			i--
		} else {
			path = append(path, 'R')
			j--
		}
	}

	fmt.Println(dp[n][m])
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Printf("%c", path[i])
	}
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
