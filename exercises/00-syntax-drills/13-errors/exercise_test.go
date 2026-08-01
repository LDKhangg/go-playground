//go:build exercise

package errvalues

import (
	"errors"
	"strings"
	"testing"
)

func TestValidateTitle(t *testing.T) {
	if err := ValidateTitle("   "); !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("expected ErrEmptyTitle, got %v", err)
	}
	if err := ValidateTitle(strings.Repeat("x", 81)); !errors.Is(err, ErrTitleTooLong) {
		t.Fatalf("expected ErrTitleTooLong, got %v", err)
	}
	if err := ValidateTitle("ship API"); err != nil {
		t.Fatalf("expected valid title, got %v", err)
	}
}
