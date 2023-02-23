package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstStr := scanner.Text()
	first := strings.Split(firstStr, " ")
	scanner.Scan()
	secondStr := scanner.Text()
	second := strings.Split(secondStr, " ")
	n := 0

	for len(first) > 0 && len(second) > 0 {
		n++
		a, b := first[0], second[0]
		first, second = first[1:], second[1:]
		aInt, _ := strconv.Atoi(a)
		bInt, _ := strconv.Atoi(b)

		if (aInt > bInt && !(b == "0" && a == "09")) || (a == "0" && b == "09") {
			first = append(first, a, b)
		} else {
			second = append(second, a, b)
		}

		if n == 1000000 {
			fmt.Println("botva")
			return
		}
	}

	if len(first) > 0 {
		fmt.Println("first", n)
	} else {
		fmt.Println("second", n)
	}
}
