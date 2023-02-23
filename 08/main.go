package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	fmt.Scan(&k)

	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64
	for i := 0; i < k; i++ {
		var x, y int
		fmt.Scan(&x, &y)

		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	fmt.Println(minX, minY, maxX, maxY)
}
