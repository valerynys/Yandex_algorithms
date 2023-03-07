package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// read n and m
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	vertex := make([]int, m)
	for i := 0; i < m; i++ {
		vertex[i] = -1
	}

	stations := make(map[int]map[int]bool)
	distance := make(map[int][]int)
	edges := make(map[int]map[int]bool)

	for i := 0; i < m; i++ {
		scanner.Scan()
		pStr := strings.Split(scanner.Text(), " ")[1:]
		p := make([]int, len(pStr))
		for j, v := range pStr {
			p[j], _ = strconv.Atoi(v)
		}
		for _, j := range p {
			if _, ok := stations[j]; !ok {
				stations[j] = make(map[int]bool)
			}
			for k := range stations[j] {
				if _, ok := edges[k]; !ok {
					edges[k] = make(map[int]bool)
				}
				edges[k][i] = true
				edges[i][k] = true
			}
			stations[j][i] = true
		}
	}

	// read a and b
	scanner.Scan()
	ab := strings.Split(scanner.Text(), " ")
	a, _ := strconv.Atoi(ab[0])
	b, _ := strconv.Atoi(ab[1])

	starts := stations[a]
	ends := stations[b]

	for i := range starts {
		vertex[i] = 0
		distance[0] = append(distance[0], i)
	}

	cur := 0
	ma := 0
	for cur <= ma {
		for _, v := range distance[cur] {
			for next := range edges[v] {
				if vertex[next] == -1 {
					vertex[next] = cur + 1
					ma = cur + 1
					distance[cur+1] = append(distance[cur+1], next)
				}
			}
		}
		cur++
	}

	minDist := -1
	for i := range ends {
		if minDist == -1 || vertex[i] < minDist {
			minDist = vertex[i]
		}
	}
	fmt.Println(minDist)
}
