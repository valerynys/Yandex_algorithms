package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i >= 2 && j >= 1 {
				dp[i][j] += dp[i-2][j-1]
			}
			if i >= 1 && j >= 2 {
				dp[i][j] += dp[i-1][j-2]
			}
		}
	}

	fmt.Println(dp[n-1][m-1])
}
