package main

import "fmt"

type Queue []int

// IsEmpty: check if 11_stack is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Push a new value onto the 11_stack
func (q *Queue) Push(i int) {
	*q = append(*q, i) // Simply append the new value to the end of the 11_stack
}

func (q *Queue) Front() int {
	return (*q)[0]
}

// Remove and return top element of 11_stack. Return false if 11_stack is empty.
func (q *Queue) Pop() int {
	element := (*q)[0] // Index into the slice and obtain the element.
	*q = (*q)[1:]      // Remove it from the 11_stack by slicing it off.
	return element
}

func (q *Queue) Size() int {
	return len(*q)
}

func main() {
	var qu Queue // create a 11_stack variable of type Stack
	var s string
	var n int
	for s != "exit" {
		fmt.Scan(&s)
		switch {
		case s == "push":
			fmt.Scan(&n)
			qu.Push(n)
			fmt.Println("ok")
		case s == "pop":
			if len(qu) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(qu.Pop())
			}
		case s == "front":
			if len(qu) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(qu.Front())
			}
		case s == "size":
			fmt.Println(qu.Size())
		case s == "clear":
			for len(qu) > 0 {
				qu.Pop()
			}
			fmt.Println("ok")
		}
	}
	fmt.Println("bye")
}
