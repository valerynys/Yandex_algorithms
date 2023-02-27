package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	cost := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cost[i])
	}

	const maxCost = 300
	const maxDays = 100
	const inf = 1e9

	// dp[i][j] - минимальная стоимость обедов для дней с 1 по i, если на текущий день есть j купонов
	var dp = [maxDays + 1][]int{}
	for i := 1; i <= maxDays; i++ {
		for j := 0; j <= maxCost*n; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= maxCost*n; j++ {
			dp[i][j] = dp[i-1][j]
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i][j-1])
			}
			if j >= cost[i-1] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-cost[i-1]]+cost[i-1])
			}
		}
	}

	var ans, k1, k2 int
	for i := 0; i <= maxCost*n; i++ {
		if dp[n][i] < ans {
			ans = dp[n][i]
			k1 = i
			k2 = maxCost*n - i
		}
	}

	fmt.Println(ans, k1)

	var days []int
	for i := n; i > 0 && k1 >= 0 && k2 >= 0; i-- {
		if dp[i][k1] == dp[i-1][k1] {
			// Не использовали купон в день i
			continue
		}
		if dp[i][k1] == dp[i][k1-1] {
			// Использовали купон в день i, но его осталось достаточно
			k1--
			k2++
			continue
		}
		// Использовали купон в день i и его осталось недостаточно
		days = append(days, i)
		k1 -= cost[i-1]
	}

	fmt.Println(k2)
	for i := len(days) - 1; i >= 0; i-- {
		fmt.Println(days[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
