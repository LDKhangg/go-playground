package tasks

import "testing"

func TestStoreAddAssignsIDAndTrimsTitle(t *testing.T) {
	store := NewStore()

	task, err := store.Add("  learn go  ")
	if err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	if task.ID != 1 {
		t.Fatalf("expected ID 1, got %d", task.ID)
	}

	if task.Title != "learn go" {
		t.Fatalf("expected trimmed title, got %q", task.Title)
	}
}

func TestStoreAddRejectsEmptyTitle(t *testing.T) {
	store := NewStore()

	_, err := store.Add("   ")
	if err != ErrEmptyTitle {
		t.Fatalf("expected ErrEmptyTitle, got %v", err)
	}
}

func TestStoreListReturnsCopy(t *testing.T) {
	store := NewStore()
	if _, err := store.Add("learn tests"); err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	list := store.List()
	list[0].Title = "changed outside"

	again := store.List()
	if again[0].Title != "learn tests" {
		t.Fatalf("expected store data to stay unchanged, got %q", again[0].Title)
	}
}
