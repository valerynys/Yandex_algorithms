package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	b := make([]int, n+1)
	c := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i], &b[i], &c[i])
	}

	t := make([]int, n+1)
	t[0] = 0
	t[1] = a[1]
	if n > 1 {
		t[2] = min(a[1]+a[2], b[1])
	}

	for i := 3; i <= n; i++ {
		t[i] = min(t[i-1]+a[i], t[i-2]+b[i-1], t[i-3]+c[i-2])
	}
	fmt.Println(t[n])
}

func min(nums ...int) int {
	smallest := nums[0]
	for _, num := range nums[1:] {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}
