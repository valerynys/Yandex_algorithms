package main

import "fmt"

func moveZeroes(nums []int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {

			continue
		}
		nums[left] = nums[i]
		left++
		fmt.Println(nums)
	}
	fmt.Println(nums)
	fmt.Println(left)
	for ; left < len(nums); left++ {
		nums[left] = 0
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{0, 3, 1, 2, 2, 3, 0, 4, 5}
	moveZeroes(nums)
}
