//go:build exercise

package functions

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	quotient, remainder, err := Divide(17, 5)
	if err != nil {
		t.Fatalf("Divide returned error: %v", err)
	}
	if quotient != 3 || remainder != 2 {
		t.Fatalf("Divide(17, 5) = (%d, %d), want (3, 2)", quotient, remainder)
	}
}

func TestDivideRejectsZeroDivisor(t *testing.T) {
	_, _, err := Divide(10, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Fatalf("Divide(10, 0) error = %v, want %v", err, ErrDivideByZero)
	}
}
