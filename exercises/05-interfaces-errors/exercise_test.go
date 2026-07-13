//go:build exercise

package interfaceserrors

import (
	"errors"
	"testing"
)

type validatorFunc func(string) error

func (f validatorFunc) Validate(title string) error { return f(title) }

func TestValidateTitleUsesValidator(t *testing.T) {
	var gotTitle string
	err := ValidateTitle(validatorFunc(func(title string) error {
		gotTitle = title
		return nil
	}), "learn interfaces")

	if err != nil {
		t.Fatalf("ValidateTitle returned error: %v", err)
	}
	if gotTitle != "learn interfaces" {
		t.Fatalf("validator received %q, want %q", gotTitle, "learn interfaces")
	}
}

func TestValidateTitleWrapsValidationError(t *testing.T) {
	err := ValidateTitle(validatorFunc(func(string) error {
		return ErrEmptyTitle
	}), "")

	if !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("ValidateTitle error = %v, want wrapped %v", err, ErrEmptyTitle)
	}
	if err.Error() == ErrEmptyTitle.Error() {
		t.Fatalf("ValidateTitle error %q has no context", err)
	}
}
