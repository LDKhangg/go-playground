package taskapp

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/LDKhangg/go-playground/apps/task-manager/internal/tasks"
)

func TestServiceListHonorsContextCancellationDuringDelay(t *testing.T) {
	store := tasks.NewStore()
	if _, err := store.Add("learn context"); err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	service := NewService(store)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	_, err := service.List(ctx, 50*time.Millisecond)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context deadline exceeded, got %v", err)
	}
}

func TestServiceCreateReturnsContextErrorBeforeWork(t *testing.T) {
	store := tasks.NewStore()
	service := NewService(store)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := service.Create(ctx, "learn context")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled, got %v", err)
	}
	if got := store.List(); len(got) != 0 {
		t.Fatalf("expected store to stay empty, got %d tasks", len(got))
	}
}
