//go:build exercise

package runes

import "testing"

func TestInitials(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "two words", in: "Mina Park", want: "MP"},
		{name: "extra spaces", in: "  Mina   Park  ", want: "MP"},
		{name: "one word", in: "Go", want: "G"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Initials(tt.in); got != tt.want {
				t.Fatalf("Initials(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
