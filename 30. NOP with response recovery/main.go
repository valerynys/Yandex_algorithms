package main

import "fmt"

const N = 1002

func lcs(x []int, y []int) []int {
	n := len(x)
	m := len(y)
	a := make([][]int, n+1)
	for i := range a {
		a[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if x[i-1] == y[j-1] {
				a[i][j] = 1 + a[i-1][j-1]
			} else {
				a[i][j] = max(a[i-1][j], a[i][j-1])
			}
		}
	}

	v := make([]int, 0)

	i := n
	j := m

	for i > 0 && j > 0 {
		if x[i-1] == y[j-1] {
			v = append(v, x[i-1])
			i--
			j--
		} else if a[i-1][j] == a[i][j] {
			i--
		} else {
			j--
		}
	}

	reverse(v)

	return v
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func reverse(x []int) {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
}

func main() {

	var n, m int
	fmt.Scan(&n)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	fmt.Scan(&m)
	y := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&y[i])
	}

	result := lcs(x, y)

	for _, val := range result {
		fmt.Printf("%d ", val)
	}

	fmt.Println()
}
