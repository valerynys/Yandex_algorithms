package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	var num int
	fmt.Scan(&num)

	res := make(map[int]int)
	ordLetter := 97
	for i := 0; i < num; i++ {
		var val int
		fmt.Scan(&val)
		res[ordLetter] = val
		ordLetter++
	}

	maxRes := 0
	for _, val := range res {
		if val > maxRes {
			maxRes = val
		}
	}

	total := 0
	node := &Node{0, nil}
	//prev := node
	for j := 0; j < maxRes; j++ {
		for key, value := range res {
			if j+1 <= value {
				if node.value == 0 {
					node.next = &Node{key, nil}
					node = node.next
				} else {
					prevValue := node.value
					if key-prevValue == 1 {
						total++
					}
					//prev = node
					node = &Node{key, nil}
				}
			}
		}
	}

	fmt.Println(total)
}
