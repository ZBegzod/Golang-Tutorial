package main

import "fmt"

func main() {
	i := 6
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Print("One")
	case 2:
		fmt.Print("Two")
	case 3:
		fmt.Print("Three")
	default:
		fmt.Print("Nothing")
	}

}
