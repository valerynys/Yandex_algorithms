package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	a := make([]int, 10001)
	dp := make([]int, 10001)

	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a[1 : n+1])

	dp[2] = a[2] - a[1]
	dp[3] = a[3] - a[1]

	for i := 4; i <= n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + a[i] - a[i-1]
	}

	fmt.Println(dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
