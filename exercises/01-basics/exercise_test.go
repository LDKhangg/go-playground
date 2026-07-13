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
		{name: "oldest child", age: 12, want: 5},
		{name: "youngest adult", age: 13, want: 12},
		{name: "adult", age: 30, want: 12},
		{name: "oldest adult", age: 64, want: 12},
		{name: "youngest senior", age: 65, want: 7},
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

func TestTotalTicketPrice(t *testing.T) {
	ages := []int{8, 30, 70}
	if got := TotalTicketPrice(ages); got != 24 {
		t.Fatalf("TotalTicketPrice(%v) = %d, want 24", ages, got)
	}
}
