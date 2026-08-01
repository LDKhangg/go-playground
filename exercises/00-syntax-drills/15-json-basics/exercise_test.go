//go:build exercise

package jsonbasics

import "testing"

func TestDecodeTask(t *testing.T) {
	t.Run("valid json", func(t *testing.T) {
		task, err := DecodeTask([]byte(`{"title":"ship API"}`))
		if err != nil {
			t.Fatalf("DecodeTask returned error: %v", err)
		}
		if task.Title != "ship API" {
			t.Fatalf("Title = %q, want %q", task.Title, "ship API")
		}
	})

	t.Run("blank title", func(t *testing.T) {
		if _, err := DecodeTask([]byte(`{"title":"   "}`)); err == nil {
			t.Fatal("expected blank title error")
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		if _, err := DecodeTask([]byte(`{"title":`)); err == nil {
			t.Fatal("expected invalid json error")
		}
	})
}
