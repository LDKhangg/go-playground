//go:build exercise

package conversion

import "testing"

func TestParsePort(t *testing.T) {
	t.Run("valid port", func(t *testing.T) {
		port, err := ParsePort("8080")
		if err != nil {
			t.Fatalf("ParsePort returned error: %v", err)
		}
		if port != 8080 {
			t.Fatalf("port = %d, want 8080", port)
		}
	})

	t.Run("invalid text", func(t *testing.T) {
		if _, err := ParsePort("abc"); err == nil {
			t.Fatal("expected an error for non-numeric input")
		}
	})

	t.Run("negative port", func(t *testing.T) {
		if _, err := ParsePort("-1"); err == nil {
			t.Fatal("expected an error for negative port")
		}
	})
}
