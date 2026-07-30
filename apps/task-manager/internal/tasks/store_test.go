package tasks

import (
	"fmt"
	"sync"
	"testing"
)

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

func TestStoreSupportsConcurrentAddAndList(t *testing.T) {
	const additions = 100

	store := NewStore()
	start := make(chan struct{})
	var workers sync.WaitGroup
	for i := range additions {
		workers.Add(2)
		go func(i int) {
			defer workers.Done()
			<-start
			if _, err := store.Add(fmt.Sprintf("task %d", i)); err != nil {
				t.Errorf("Add returned error: %v", err)
			}
		}(i)
		go func() {
			defer workers.Done()
			<-start
			_ = store.List()
		}()
	}

	close(start)
	workers.Wait()

	listed := store.List()
	if len(listed) != additions {
		t.Fatalf("List returned %d tasks after %d successful additions", len(listed), additions)
	}

	seenIDs := make(map[int]bool, additions)
	for _, task := range listed {
		if task.ID < 1 || task.ID > additions {
			t.Fatalf("task ID = %d, want a value from 1 through %d", task.ID, additions)
		}
		if seenIDs[task.ID] {
			t.Fatalf("task ID %d was assigned more than once", task.ID)
		}
		seenIDs[task.ID] = true
	}
}
