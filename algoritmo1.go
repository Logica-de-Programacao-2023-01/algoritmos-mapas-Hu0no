package main

import (
	"fmt"
	"strings"
)

func CountWords(x string) map[string]int {
	wordCount := make(map[string]int)
	x = strings.ToLower(x)
	words := strings.Fields(x)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}
func main() {
	text := "Olá, este é um exemplo de frase. Este é um exemplo simples."
	wordCount := CountWords(text)
	for word, count := range wordCount {
		fmt.Printf("Palavra: %s\tContagem: %d\n", word, count)
	}
}
