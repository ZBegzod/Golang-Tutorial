package main

import "fmt"

func update(menu map[string]float64) {
	menu["pia"] = 5.55
	menu["coffee"] = 1.55

}
func main() {
	menu := map[string]float64{
		"soup":   5.40,
		"salad":  4.40,
		"pia":    3.40,
		"toffee": 2.55,
	}
	//menu["pia"] = 2.8

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	update(menu)
	fmt.Println(menu)
}
