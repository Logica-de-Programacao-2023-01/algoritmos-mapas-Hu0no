package main

import (
	"fmt"
	"unicode"
)

func FrequencyOfLetters(x string) map[rune]int {
	frequency := make(map[rune]int)
	for _, letter := range x {
		letter = unicode.ToLower(letter)
		if !unicode.IsLetter(letter) {
			continue
		}
		frequency[letter]++
	}
	return frequency
}
func main() {
	str := "Hello, World!"

	letterFrequency := FrequencyOfLetters(str)

	for char, count := range letterFrequency {
		fmt.Printf("Letra: %c - Frequência: %d\n", char, count)
	}
}
