package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	sectors := make([]bool, m+1)
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		for j := a; j <= b; j++ {
			if sectors[j] {
				// если сектор уже был занят, то предыдущий раздел затирается
				sectors = make([]bool, m+1)
				break
			}
			sectors[j] = true
		}
	}

	// считаем количество работающих операционных систем,
	// исключая секторы, которые не были заняты ни разу
	count := 0
	for i := 1; i <= m; i++ {
		if sectors[i] {
			count++
		}
	}
	fmt.Println(count)
}
