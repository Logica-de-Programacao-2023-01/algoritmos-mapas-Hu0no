package main

import "fmt"

func FibonnaciSequence(n int) map[int]int {
	x := make(map[int]int)
	x[0] = 0
	x[1] = 1
	for y := 2; y < n; y++ {
		x[y] = x[y-2] + x[y-1]
	}
	return x
}
func main() {
	n := 8
	fmt.Print(FibonnaciSequence(n))
}
