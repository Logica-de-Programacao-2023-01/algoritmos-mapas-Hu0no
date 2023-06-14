package main

import "fmt"

func SumOfFrequencies(y []map[string]int) map[string]int {
	x := make(map[string]int)
	for _, wordCounts := range y {
		for word, count := range wordCounts {
			x[word] += count
		}
	}

	return x
}

func main() {
	wordCountsList := []map[string]int{
		{"hello": 2, "world": 1, "example": 3},
		{"hello": 1, "world": 2},
		{"hello": 3, "example": 2},
	}

	mergedCounts := SumOfFrequencies(wordCountsList)

	fmt.Println(mergedCounts)
}
