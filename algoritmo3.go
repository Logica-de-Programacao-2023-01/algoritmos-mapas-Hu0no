package main

import "fmt"

func AddValues(x map[string]int) int {
	Sum := 0
	for _, value := range x {
		Sum += value
	}
	return Sum
}
func main() {
	x := map[string]int{
		"Crianças": 50,
		"Adultos":  200,
		"Idosos":   70,
	}
	fmt.Print(AddValues(x))
}
