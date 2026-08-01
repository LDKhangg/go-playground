//go:build exercise

package zerovalues

import "testing"

func TestDefaultProfileReturnsZeroValues(t *testing.T) {
	profile := DefaultProfile()
	if profile.Name != "" {
		t.Fatalf("Name = %q, want empty string", profile.Name)
	}
	if profile.Active {
		t.Fatalf("Active = true, want false")
	}
	if profile.Attempts != 0 {
		t.Fatalf("Attempts = %d, want 0", profile.Attempts)
	}
}
