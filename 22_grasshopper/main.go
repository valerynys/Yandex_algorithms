package main

import "fmt"

func main() {
	var i, n, k int
	dp := make([]int, 35)
	fmt.Scan(&n, &k)
	dp[1], dp[2] = 1, 1
	for i = 3; i <= k; i++ {
		dp[i] = 2 * dp[i-1]
	}
	for ; i <= n; i++ {
		dp[i] = 2*dp[i-1] - dp[i-k-1]
	}
	fmt.Println(dp[n])
}
