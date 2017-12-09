package exercise

import (
	"strings"
	"unicode"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	data := make(map[string]int)

	for _, word := range words {
		data[word] = len(word)
	}

	return data
}

func IsNotSpace(r rune) bool {
	return !unicode.IsSpace(r)
}

func Trim(s string) string {
	first := strings.IndexFunc(s, IsNotSpace)
	last := strings.LastIndexFunc(s, IsNotSpace)

	return s[first:last+1]
}