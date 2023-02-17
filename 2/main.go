package main

import (
	"fmt"
)

func maxBeauty(s string, k int) int {
	n := len(s)
	count := make([]int, 26)
	ans := 0
	maxCount := 0
	left := 0
	for right := 0; right < n; right++ {
		count[int(s[right]-'a')]++
		maxCount = max(maxCount, count[int(s[right]-'a')])
		if right-left+1-maxCount > k {
			count[int(s[left]-'a')]--
			left++
		}
		ans = max(ans, min(right-left+1, maxCount+k))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var k int
	var s string
	fmt.Scan(&k)
	fmt.Scan(&s)
	fmt.Println(maxBeauty(s, k))
}
