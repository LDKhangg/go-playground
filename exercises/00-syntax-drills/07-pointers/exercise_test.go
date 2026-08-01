//go:build exercise

package pointers

import "testing"

func TestIncrement(t *testing.T) {
	value := 4
	if err := Increment(&value); err != nil {
		t.Fatalf("Increment returned error: %v", err)
	}
	if value != 5 {
		t.Fatalf("value = %d, want 5", value)
	}
}

func TestIncrementRejectsNil(t *testing.T) {
	if err := Increment(nil); err == nil {
		t.Fatal("expected nil pointer error")
	}
}
