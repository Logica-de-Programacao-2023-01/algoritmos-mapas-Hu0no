package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	// Converte a string para uma fatia de caracteres
	characters := strings.Split(s, "")
	// Ordena a fatia de caracteres em ordem alfabética
	sort.Strings(characters)
	// Junta os caracteres ordenados para formar uma string
	return strings.Join(characters, "")
}

func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		// Ordena as letras da palavra atual em ordem alfabética
		sortedWord := sortString(word)

		// Adiciona a palavra ao grupo de anagramas correspondente
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	return anagramMap
}

func main() {
	words := []string{"listen", "silent", "elbow", "below", "state", "taste"}

	anagramGroups := findAnagrams(words)

	for _, group := range anagramGroups {
		fmt.Println(group)
	}
}
