package main

import (
	"fmt"
	"strconv"
	"strings"
)

func evaluatePostfixExpression(expression string) int {
	tokens := strings.Split(expression, " ")
	stack := []int{}
	for _, token := range tokens {
		if token == "+" {
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := op1 + op2
			stack = append(stack, result)
		} else if token == "-" {
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := op1 - op2
			stack = append(stack, result)
		} else if token == "*" {
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			result := op1 * op2
			stack = append(stack, result)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func main() {
	var expression string
	fmt.Scan(&expression)
	fmt.Println(evaluatePostfixExpression(expression))
}
