package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
	z int
}

func main() {
	// read input
	var n int
	fmt.Scan(&n)

	cave := make([][][]byte, n)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan() // skip empty line
		cave[i] = make([][]byte, n)

		for j := 0; j < n; j++ {
			scanner.Scan()
			cave[i][j] = []byte(scanner.Text())
		}
	}

	// find starting position and initialize BFS queue
	var start_pos position
	for z := 0; z < n; z++ {
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				if cave[z][y][x] == 'S' {
					start_pos = position{x, y, z}
					break
				}
			}
		}
	}

	queue := make([][2]interface{}, 0)
	queue = append(queue, [2]interface{}{start_pos, 0})

	// initialize visited set
	visited := make(map[position]bool)
	visited[start_pos] = true

	// BFS loop
	for len(queue) > 0 {
		pos := queue[0][0].(position)
		dist := queue[0][1].(int)
		queue = queue[1:]

		x, y, z := pos.x, pos.y, pos.z

		// check if we've reached the surface
		if z == 0 {
			fmt.Println(dist)
			break
		}

		// generate all possible moves from current position
		moves := []position{
			{x + 1, y, z},
			{x - 1, y, z},
			{x, y + 1, z},
			{x, y - 1, z},
			{x, y, z + 1},
			{x, y, z - 1},
		}

		// process each move
		for _, move := range moves {
			x2, y2, z2 := move.x, move.y, move.z
			if 0 <= x2 && x2 < n && 0 <= y2 && y2 < n && 0 <= z2 && z2 < n { // check if move is within cave bounds
				if cave[z2][y2][x2] == '.' && !visited[move] { // check if move is valid
					queue = append(queue, [2]interface{}{move, dist + 1})
					visited[move] = true
				}
			}
		}
	}
}
