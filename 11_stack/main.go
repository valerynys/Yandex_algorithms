package main

import "fmt"

type Stack []int

// IsEmpty: check if 11_stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the 11_stack
func (s *Stack) Push(i int) {
	*s = append(*s, i) // Simply append the new value to the end of the 11_stack
}

func (s *Stack) Back() int {
	index := len(*s) - 1 // Get the index of the top most element.
	return (*s)[index]
}

// Remove and return top element of 11_stack. Return false if 11_stack is empty.
func (s *Stack) Pop() int {
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the 11_stack by slicing it off.
	return element
}

func (s *Stack) Size() int {
	return len(*s)
}

func main() {
	var stack Stack // create a 11_stack variable of type Stack
	var s string
	var n int
	for s != "exit" {
		fmt.Scan(&s)
		switch {
		case s == "push":
			fmt.Scan(&n)
			stack.Push(n)
			fmt.Println("ok")
		case s == "pop":
			if len(stack) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(stack.Pop())
			}
		case s == "back":
			if len(stack) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(stack.Back())
			}
		case s == "size":
			fmt.Println(stack.Size())
		case s == "clear":
			for len(stack) > 0 {
				stack.Pop()
			}
			fmt.Println("ok")
		}
	}
	fmt.Println("bye")
}
