package main

import (
	"fmt"
	"strings"
)

var numberOfList []int

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(s []string, f func(string)) {
	for _, value := range s {
		f(value)
	}
}
func calculator(symbol string, a int, b int) int {
	if symbol == "+" {
		return a + b
	} else if symbol == "-" {
		return a - b
	} else if symbol == "/" {
		if a != 0 && b != 0 {
			return a / b
		}
	} else if symbol == "*" {
		return a * b
	}

	return a
}
func getInitialize(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, "")
	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	//cycleNames([]string{"Sardor", "Aziz", "Davron"}, sayGreeting)
	//result := calculator("*", 5, 6)
	//fmt.Println(result)

	a, b := getInitialize("Golang Language")
	fmt.Println(a, b)

}
