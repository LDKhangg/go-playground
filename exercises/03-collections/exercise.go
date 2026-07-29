package collections

func UniqueWords(words []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, word := range words {
		if !seen[word] {
			seen[word] = true
			result = append(result, word)
		}
	}

	return result
}

func SumArray(numbers [4]int) int {
	var total int
	for i := 0; i < 4; i++ {
		total += numbers[i]
	}
	return total
}
