//go:build exercise

package branching

import "testing"

func TestClassifyScore(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{score: -1, want: "invalid"},
		{score: 40, want: "fail"},
		{score: 70, want: "pass"},
		{score: 90, want: "excellent"},
		{score: 120, want: "invalid"},
	}

	for _, tt := range tests {
		if got := ClassifyScore(tt.score); got != tt.want {
			t.Fatalf("ClassifyScore(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}
