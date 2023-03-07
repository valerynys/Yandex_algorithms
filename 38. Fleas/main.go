package main

import (
	"fmt"
)

type pair struct {
	x, y int
}

func getNeighbours(x, y, n, m int) []pair {
	neighbours := make([]pair, 0)
	for _, i := range []pair{{x - 2, y - 1}, {x - 2, y + 1}, {x - 1, y - 2}, {x - 1, y + 2}, {x + 1, y - 2}, {x + 1, y + 2}, {x + 2, y - 1}, {x + 2, y + 1}} {
		if i.x >= 0 && i.x < n && i.y >= 0 && i.y < m {
			neighbours = append(neighbours, i)
		}
	}
	return neighbours
}

func main() {
	var n, m, s, t, q int
	fmt.Scan(&n, &m, &s, &t, &q)
	s--
	t--

	distance := make(map[int][]pair)
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, m)
		for j := range mat[i] {
			mat[i][j] = -1
		}
	}

	mat[s][t] = 0
	distance[0] = append(distance[0], pair{s, t})

	cur := 0
	ma := 0
	for cur <= ma {
		for _, v := range distance[cur] {
			for _, next := range getNeighbours(v.x, v.y, n, m) {
				if mat[next.x][next.y] == -1 {
					mat[next.x][next.y] = cur + 1
					distance[cur+1] = append(distance[cur+1], next)
					ma = cur + 1
				}
			}
		}
		cur++
	}

	sol := 0
	for i := 0; i < q; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		if mat[a][b] == -1 {
			fmt.Println(-1)
			sol = -1
			break
		}
		sol += mat[a][b]
	}

	if sol >= 0 {
		fmt.Println(sol)
	}
}
