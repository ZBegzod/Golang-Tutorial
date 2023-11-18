package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{54, 48, 87}
	//var ages = [3]int{54, 48, 87}

	fruits := []string{"apple", "banana", "orange", "liman"}
	//fmt.Println(fruits, len(fruits))

	// slice (use arrays under the hood)
	//var scores = []int{100, 56, 89}
	//scores[0] = 54
	//scores = append(scores, 45)
	//fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := fruits[1:3]
	//rangeTwo := fruits[1:]
	//rangeThree := fruits[:2]
	rangeOne = append(rangeOne, "juice", "asdas")
	fmt.Println(rangeOne)
}
