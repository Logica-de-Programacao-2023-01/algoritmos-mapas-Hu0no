package main

import "fmt"

func AddMaps(x map[string]int, y map[string]int) map[string]int {
	z := make(map[string]int)
	for key, value := range x {
		z[key] = value
	}
	for key, value := range y {
		z[key] = value
	}
	return z
}
func main() {
	x := map[string]int{
		"Banana": 5,
		"Maçã":   8,
	}
	y := map[string]int{
		"Pêra":  3,
		"Limão": 2,
		"Maçã":  4,
	}
	fmt.Print(AddMaps(x, y))
}
