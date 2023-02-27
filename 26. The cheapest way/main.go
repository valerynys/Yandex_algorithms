package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
		for j := 0; j < M; j++ {
			dp[i][j] = 1e9
			if i == 0 && j == 0 {
				dp[i][j] = a[i][j]
			} else {
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j])
				}
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1])
				}
				dp[i][j] += a[i][j]
			}
		}
	}

	fmt.Println(dp[N-1][M-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
