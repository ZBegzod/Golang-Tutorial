package main

import (
	"fmt"
	"sort"
)

func main() {
	//greeting := "hello world friends"

	// Contains metodi string tarkibida belgilardan biriga mos kelsa true qiymatini qaytaradi
	//fmt.Println(strings.Contains(greeting, "hello"))

	// ReplaceAll - stringdan berilgan elementni almashtiradi
	//fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))

	// ToUpper - berilgan stringni katta harflarga o'girib beradi
	//fmt.Println(strings.ToUpper(greeting))

	// Index - berilgan belgining indexini qaytaradi
	//fmt.Println(strings.Index(greeting, "f"))

	// Split - berilgan stringni listga o'tkazadi
	//res := strings.Split(greeting, " ")
	//fmt.Println(res)

	ages := []int{10: 5, 1, 3, 2: 5}
	//sort.Ints(ages)
	//fmt.Println(ages)
	//index := sort.SearchInts(ages, 5)
	//fmt.Println(index)

	fruits := []string{"apple", "banana", "orange", "kiwi"}
	sort.Strings(fruits)
	fmt.Println(fruits)
	index := sort.SearchStrings(fruits, "orange")
	fmt.Println(index)

	fmt.Println("original string value =", ages)
	fmt.Println("length: ", len(ages))
}
