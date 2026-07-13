//go:build exercise

package collections

import (
	"slices"
	"testing"
)

func TestUniqueWordsPreservesFirstOccurrenceOrder(t *testing.T) {
	input := []string{"go", "test", "go", "maps", "test"}
	want := []string{"go", "test", "maps"}

	got := UniqueWords(input)
	if !slices.Equal(got, want) {
		t.Fatalf("UniqueWords(%v) = %v, want %v", input, got, want)
	}
}

func TestUniqueWordsDoesNotReuseInputStorage(t *testing.T) {
	input := []string{"go", "test"}
	got := UniqueWords(input)
	if len(got) != len(input) {
		t.Fatalf("UniqueWords(%v) length = %d, want %d", input, len(got), len(input))
	}
	got[0] = "changed"

	if input[0] != "go" {
		t.Fatalf("UniqueWords modified or reused input storage: input = %v", input)
	}
}
