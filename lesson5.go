package main

import "fmt"

var a = 4
var b = 8
var c = 12
var number = 458

func main() {

	//if a > b && a > c {
	//	fmt.Printf("max: %v", a)
	//} else if b > a && b > c {
	//	fmt.Printf("max: %v", b)
	//} else if c > a && c > b {
	//	fmt.Printf("max: %v", c)
	//}

	//if number >= 1 && number <= 9 {
	//	fmt.Print("Bir honali ")
	//	if number%2 == 0 {
	//		fmt.Print("Juft son")
	//	} else {
	//		fmt.Print("Toq son")
	//	}
	//} else if number >= 10 && number <= 99 {
	//	fmt.Print("Ikki honali ")
	//	if number%2 == 0 {
	//		fmt.Print("Juft son")
	//	} else {
	//		fmt.Print("Toq son")
	//	}
	//} else if number >= 100 && number <= 999 {
	//	fmt.Print("Uch honali ")
	//	if number%2 == 0 {
	//		fmt.Print("Juft son")
	//	} else {
	//		fmt.Print("Toq son")
	//	}
	//}

	fruits := []string{"apple", "banana", "orange", "liman"}
	//for i := 0; i < len(fruits); i++ {
	//	fmt.Println(fruits[i])
	//}
	for index, value := range fruits {
		//fmt.Println(fmt.Sprintf("%v %v", index, value))
		if index%2 == 0 {
			fmt.Println("continuing at pos", index)
			continue
		}
		fmt.Printf("the value at index %v is %v \n", index, value)
	}
}
