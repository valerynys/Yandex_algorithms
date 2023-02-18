package main

import "fmt"

func heapSort(arr []int) {
	// Построение Max Heap
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	// Извлечение максимального элемента и перестройка дерева
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func heapify(arr []int, i, n int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, largest, n)
	}
}

func read(n int) []int {
	in := make([]int, n)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			return in[:i]
		}
	}
	return in
}

func main() {
	var ar []int
	var n int
	fmt.Scan(&n)
	ar = read(n)
	heapSort(ar)
	for i := 0; i < n; i++ {
		fmt.Print(ar[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
}
