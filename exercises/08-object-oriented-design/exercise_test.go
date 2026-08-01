//go:build exercise

package oopdesign

import (
	"errors"
	"testing"
	"time"
)

type stubClock struct {
	now time.Time
}

func (c stubClock) Now() time.Time {
	return c.now
}

type stubIDs struct {
	next string
}

func (g stubIDs) NextID() string {
	return g.next
}

type fakeRepository struct {
	tasks      map[string]Task
	savedTasks []Task
	saveErr    error
	findErr    error
}

func newFakeRepository() *fakeRepository {
	return &fakeRepository{tasks: make(map[string]Task)}
}

func (r *fakeRepository) Save(task Task) error {
	if r.saveErr != nil {
		return r.saveErr
	}
	r.savedTasks = append(r.savedTasks, task)
	r.tasks[task.ID] = task
	return nil
}

func (r *fakeRepository) FindByID(id string) (Task, error) {
	if r.findErr != nil {
		return Task{}, r.findErr
	}
	task, ok := r.tasks[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return task, nil
}

func TestTaskServiceCreateUsesInjectedDependencies(t *testing.T) {
	now := time.Date(2026, time.August, 1, 9, 0, 0, 0, time.UTC)
	repo := newFakeRepository()
	service := NewTaskService(repo, stubClock{now: now}, stubIDs{next: "task-42"})

	task, err := service.Create("  learn service boundaries  ")
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}

	if task.ID != "task-42" {
		t.Fatalf("expected injected ID, got %q", task.ID)
	}
	if task.Title != "learn service boundaries" {
		t.Fatalf("expected trimmed title, got %q", task.Title)
	}
	if !task.CreatedAt.Equal(now) {
		t.Fatalf("expected created time %v, got %v", now, task.CreatedAt)
	}
	if len(repo.savedTasks) != 1 {
		t.Fatalf("expected repository save call, got %d", len(repo.savedTasks))
	}
	if repo.savedTasks[0] != task {
		t.Fatalf("expected saved task %+v, got %+v", task, repo.savedTasks[0])
	}
}

func TestTaskServiceCreateRejectsEmptyTitle(t *testing.T) {
	service := NewTaskService(newFakeRepository(), stubClock{}, stubIDs{next: "task-1"})

	_, err := service.Create("   ")
	if !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("expected ErrEmptyTitle, got %v", err)
	}
}

func TestTaskMarkDoneSetsDoneAndCompletionTime(t *testing.T) {
	now := time.Date(2026, time.August, 1, 10, 30, 0, 0, time.UTC)
	task := Task{ID: "task-7", Title: "ship feature"}

	task.MarkDone(now)

	if !task.Done {
		t.Fatal("expected task to be marked done")
	}
	if !task.CompletedAt.Equal(now) {
		t.Fatalf("expected completion time %v, got %v", now, task.CompletedAt)
	}
}

func TestTaskServiceCompleteMarksAndPersistsTask(t *testing.T) {
	now := time.Date(2026, time.August, 1, 11, 45, 0, 0, time.UTC)
	repo := newFakeRepository()
	repo.tasks["task-9"] = Task{ID: "task-9", Title: "write tests"}
	service := NewTaskService(repo, stubClock{now: now}, stubIDs{next: "unused"})

	updated, err := service.Complete("task-9")
	if err != nil {
		t.Fatalf("Complete returned error: %v", err)
	}

	if !updated.Done {
		t.Fatal("expected completed task to be done")
	}
	if !updated.CompletedAt.Equal(now) {
		t.Fatalf("expected completion time %v, got %v", now, updated.CompletedAt)
	}
	if len(repo.savedTasks) != 1 {
		t.Fatalf("expected one save after completion, got %d", len(repo.savedTasks))
	}
	if repo.savedTasks[0] != updated {
		t.Fatalf("expected saved updated task %+v, got %+v", updated, repo.savedTasks[0])
	}
}

func TestTaskServiceCompleteReturnsNotFound(t *testing.T) {
	service := NewTaskService(newFakeRepository(), stubClock{}, stubIDs{next: "unused"})

	_, err := service.Complete("missing")
	if !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
}
