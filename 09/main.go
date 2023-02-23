package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
		for j := range matrix[i] {
			fmt.Scan(&matrix[i][j])
		}
	}

	cumSum := make([][]int, n+1)
	for i := range cumSum {
		cumSum[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			cumSum[i][j] = matrix[i-1][j-1] + cumSum[i-1][j] + cumSum[i][j-1] - cumSum[i-1][j-1]
		}
	}

	results := make([]int, k)
	for i := 0; i < k; i++ {
		var x1, y1, x2, y2 int
		fmt.Scan(&x1, &y1, &x2, &y2)

		sum := cumSum[x2][y2] - cumSum[x1-1][y2] - cumSum[x2][y1-1] + cumSum[x1-1][y1-1]
		results[i] = sum
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
