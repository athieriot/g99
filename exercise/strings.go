package exercise

import "strings"

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	data := make(map[string]int)

	for _, word := range words {
		data[word] = len(word)
	}

	return data
}