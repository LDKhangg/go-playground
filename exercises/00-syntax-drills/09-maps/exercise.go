package mapsx

import "strings"

func CountWords(text string) map[string]int {
	words := strings.Fields(text)
	res := make(map[string]int)
	for _, word := range words {
		res[word]++
	}
	return res
}

func LookupScore(scores map[string]int, name string) (int, bool) {
	value, ok := scores[name]

	return value, ok
}
