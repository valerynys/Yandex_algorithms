package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Scan(&word)

	freq := make(map[rune]int)

	for i := 2; i < len(word); i++ {
		for j := i; j < len(word)-1; j++ {
			freq[rune(word[j])]++
		}
	}

	for letter := 'a'; letter <= 'z'; letter++ {
		if freq[letter] > 0 {
			fmt.Printf("%c: %d\n", letter, freq[letter])
		}
	}
}
