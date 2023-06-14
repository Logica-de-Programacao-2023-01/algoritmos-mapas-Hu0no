package main

import (
	"fmt"
	"strings"
)

func countLetters(word string) map[rune]int {
	letterCount := make(map[rune]int)

	for _, char := range word {
		letterCount[char]++
	}

	return letterCount
}

func countLettersInWords(sentence string) map[string]map[rune]int {
	words := strings.Fields(sentence)
	letterCounts := make(map[string]map[rune]int)

	for _, word := range words {
		letterCounts[word] = countLetters(word)
	}

	return letterCounts
}

func main() {
	sentence := "Olá meu caro amigo"

	letterCounts := countLettersInWords(sentence)

	for word, count := range letterCounts {
		fmt.Printf("Palavra: %s\n", word)
		for char, charCount := range count {
			fmt.Printf("Letra: %c - Contagem: %d\n", char, charCount)
		}
		fmt.Println()
	}
}
