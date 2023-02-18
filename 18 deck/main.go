package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	deque := list.New()
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			fmt.Println("bye")
			return
		} else if input == "size" {
			fmt.Println(deque.Len())
		} else if input == "clear" {
			deque.Init()
			fmt.Println("ok")
		} else if strings.HasPrefix(input, "push_front") {
			value, err := strconv.Atoi(strings.TrimPrefix(input, "push_front "))
			if err != nil {
				fmt.Println("error")
			} else {
				deque.PushFront(value)
				fmt.Println("ok")
			}
		} else if strings.HasPrefix(input, "push_back") {
			value, err := strconv.Atoi(strings.TrimPrefix(input, "push_back "))
			if err != nil {
				fmt.Println("error")
			} else {
				deque.PushBack(value)
				fmt.Println("ok")
			}
		} else if input == "pop_front" {
			if deque.Len() == 0 {
				fmt.Println("error")
			} else {
				value := deque.Remove(deque.Front())
				fmt.Println(value)
			}
		} else if input == "pop_back" {
			if deque.Len() == 0 {
				fmt.Println("error")
			} else {
				value := deque.Remove(deque.Back())
				fmt.Println(value)
			}
		} else if input == "front" {
			if deque.Len() == 0 {
				fmt.Println("error")
			} else {
				value := deque.Front().Value.(int)
				fmt.Println(value)
			}
		} else if input == "back" {
			if deque.Len() == 0 {
				fmt.Println("error")
			} else {
				value := deque.Back().Value.(int)
				fmt.Println(value)
			}
		} else {
			fmt.Println("error")
		}
	}
}
