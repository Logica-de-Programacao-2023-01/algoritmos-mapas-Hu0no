package main

import "fmt"

func FindPairNumbers(x []int) map[int]int {
	y := make(map[int]int)
	for _, number := range x {
		if number%2 == 0 {
			y[number]++
		}
	}
	return y
}
func main() {
	x := []int{2, 3, 4, 5, 6, 7, 8, 2, 4, 1}
	fmt.Print(FindPairNumbers(x))
}
