package main

import (
	"bufio"
	"fmt"
	"os"
)

// Struct to represent a point on the board
type Point struct {
	X int
	Y int
}

// Function to check if a point is on the board
func isValidMove(x, y, n, m int) bool {
	return x >= 1 && x <= n && y >= 1 && y <= m
}

// Function to calculate the distance from point (x, y) to point (targetX, targetY)
func bfs(x, y, targetX, targetY, n, m int) int {
	// Initialize the distance array and queue for traversal
	dist := make([][]int, n+1)
	for i := range dist {
		dist[i] = make([]int, m+1)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[x][y] = 0
	q := []Point{{x, y}}

	// Breadth-first search
	for len(q) > 0 {
		currX, currY := q[0].X, q[0].Y
		q = q[1:]
		for _, move := range []Point{{2, 1}, {1, 2}, {-2, 1}, {-1, 2}, {2, -1}, {1, -2}, {-2, -1}, {-1, -2}} {
			nextX, nextY := currX+move.X, currY+move.Y
			if isValidMove(nextX, nextY, n, m) && dist[nextX][nextY] == -1 {
				dist[nextX][nextY] = dist[currX][currY] + 1
				q = append(q, Point{nextX, nextY})
			}
		}
	}

	return dist[targetX][targetY]
}

func main() {
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	n, m, s, t, q := 0, 0, 0, 0, 0
	fmt.Scan(&n, &m, &s, &t, &q)
	blobs := make([]Point, q)
	for i := 0; i < q; i++ {
		x, y := 0, 0
		fmt.Scan(&x, &y)
		blobs[i] = Point{x, y}
	}

	// Calculate the distance to the feeding trough for each blob
	distances := make([]int, q)
	for i, blob := range blobs {
		d := bfs(blob.X, blob.Y, s, t, n, m)
		if d == -1 {
			fmt.Println("-1")
			return
		}
		distances[i] = d
	}

	// If all blobs can reach the feeding trough, print the sum of their distances
	sum := 0
	for _, dist := range distances {
		sum += dist
	}
	fmt.Println(sum)
}
