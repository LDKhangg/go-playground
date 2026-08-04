package runes

import "strings"

func Initials(name string) string {
	words := strings.Fields(name)
	var sb strings.Builder

	for _, word := range words {
		for _, char := range word {
			sb.WriteRune(char)
			break
		}
	}
	return sb.String()
}
