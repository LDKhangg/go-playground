//go:build exercise

package basics

import "testing"

func TestTicketPrice(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want int
	}{
		{name: "child", age: 8, want: 5},
		{name: "adult", age: 30, want: 12},
		{name: "senior", age: 70, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TicketPrice(tt.age); got != tt.want {
				t.Fatalf("TicketPrice(%d) = %d, want %d", tt.age, got, tt.want)
			}
		})
	}
}
