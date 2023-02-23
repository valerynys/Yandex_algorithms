package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	expected := 1
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		var car int
		fmt.Scan(&car)
		stack = append(stack, car)
		for len(stack) > 0 && stack[len(stack)-1] == expected {
			stack = stack[:len(stack)-1]
			expected++
		}
	}
	if len(stack) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
