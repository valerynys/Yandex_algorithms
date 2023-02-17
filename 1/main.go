package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	// Read input text
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	text := scanner.Text()

	// Count the frequency of each character
	freq := make(map[rune]int)
	for _, c := range text {
		if c != ' ' && c != '\n' {
			freq[c]++
		}
	}

	// Sort the characters in increasing order of their Unicode codes
	chars := make([]rune, 0, len(freq))
	for c := range freq {
		chars = append(chars, c)
	}
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

	// Determine the maximum frequency
	maxFreq := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
		}
	}

	// Create a string for each row of the histogram
	rows := make([]string, maxFreq)
	for i := range rows {
		row := ""
		for _, c := range chars {
			if freq[c] >= maxFreq-i {
				row += "#"
			} else {
				row += " "
			}
		}
		rows[i] = row
	}

	// Join the rows and add the character labels at the bottom
	output := strings.Join(rows, "\n") + "\n" + string(chars)

	// Print the output
	fmt.Println(output)
}
