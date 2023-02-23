package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	dp := make([]int, 40)
	dp[1], dp[2], dp[3] = 2, 4, 7

	i := 4
	for i <= n {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
		i++
	}

	fmt.Println(dp[n])
}
