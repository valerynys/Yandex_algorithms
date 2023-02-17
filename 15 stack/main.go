package main

import (
	"fmt"
)

func main() {
	var nums []int
	var stck []int
	var result []int

	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	result = make([]int, n)
	for i := range result {
		result[i] = -1
	}

	for i := 0; i < len(nums); i++ {
		for len(stck) > 0 && nums[stck[len(stck)-1]] > nums[i] {
			result[stck[len(stck)-1]] = i
			stck = stck[:len(stck)-1]
		}
		stck = append(stck, i)
	}

	for i := 0; i < len(result); i++ {
		fmt.Printf("%d ", result[i])
	}
}
