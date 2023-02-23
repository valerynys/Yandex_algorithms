package main

import (
	"fmt"
)

func f(n int) (int, []int) {
	cache := make([]int, n+1)
	prev := make([]int, n+1)

	for i := 2; i <= n; i++ {
		v := cache[i-1]
		p := i - 1
		if i%2 == 0 {
			if cache[i/2] < v {
				v = cache[i/2]
				p = i / 2
			}
		}
		if i%3 == 0 {
			if cache[i/3] < v {
				v = cache[i/3]
				p = i / 3
			}
		}
		cache[i] = v + 1
		prev[i] = p
	}
	seq := []int{n}
	for prev[seq[len(seq)-1]] != -1 {
		seq = append(seq, prev[seq[len(seq)-1]])
	}
	reverse(seq)
	return cache[n], seq
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	var num int
	num = 5
	ops, seq := f(num)
	fmt.Println(ops)
	fmt.Println(seq)
}
