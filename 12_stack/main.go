package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[', '{', '(':
			stack = append(stack, s[i])
		case ']', '}', ')':
			if len(stack) == 0 {
				return false
			}
			ch := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (s[i] == ']' && ch != '[') || (s[i] == '}' && ch != '{') || (s[i] == ')' && ch != '(') {
				return false
			}
		}
	}
	return len(stack) == 0
}
func main() {
	var s string
	fmt.Scan(&s)
	if isValid(s) == true {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
