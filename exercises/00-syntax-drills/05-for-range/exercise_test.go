//go:build exercise

package loops

import "testing"

func TestSumUntil(t *testing.T) {
	if got := SumUntil([]int{3, 4, 99, 10}, 99); got != 7 {
		t.Fatalf("SumUntil stopped total = %d, want 7", got)
	}
	if got := SumUntil(nil, 99); got != 0 {
		t.Fatalf("SumUntil(nil, 99) = %d, want 0", got)
	}
}
