package main

import "fmt"

var numbers = []int{}

func main() {
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			numbers = append(numbers, i)
		}
		i += 1
	}
	fmt.Printf("List of numbers: %v", numbers)
}
