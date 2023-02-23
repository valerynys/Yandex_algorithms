package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Heap struct {
	data []int
}

func (h *Heap) Insert(k int) {
	h.data = append(h.data, k)
	h.siftUp(len(h.data) - 1)
}

func (h *Heap) Extract() int {
	n := len(h.data)
	max := h.data[0]
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	h.siftDown(0)
	return max
}

func (h *Heap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent] >= h.data[i] {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *Heap) siftDown(i int) {
	n := len(h.data)
	for i < n {
		left := 2*i + 1
		right := 2*i + 2
		max := i
		if left < n && h.data[left] > h.data[max] {
			max = left
		}
		if right < n && h.data[right] > h.data[max] {
			max = right
		}
		if max == i {
			break
		}
		h.data[max], h.data[i] = h.data[i], h.data[max]
		i = max
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	n := readInt(scanner)
	h := &Heap{}
	for i := 0; i < n; i++ {
		op := readInt(scanner)
		if op == 0 {
			k := readInt(scanner)
			h.Insert(k)
		} else {
			fmt.Println(h.Extract())
		}
	}
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
