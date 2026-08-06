//go:build exercise

package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type memoryStore struct {
	tasks  map[string]Task
	nextID int
}

func newMemoryStore() *memoryStore {
	return &memoryStore{tasks: make(map[string]Task), nextID: 1}
}

func (s *memoryStore) Create(task Task) (Task, error) {
	id := fmt.Sprintf("task-%d", s.nextID)
	s.nextID++
	task.ID = id
	s.tasks[id] = task
	return task, nil
}

func (s *memoryStore) List() ([]Task, error) {
	out := make([]Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		out = append(out, task)
	}
	return out, nil
}

func (s *memoryStore) FindByID(id string) (Task, error) {
	task, ok := s.tasks[id]
	if !ok {
		return Task{}, ErrNotFound
	}
	return task, nil
}

func TestCreateTask(t *testing.T) {
	store := newMemoryStore()
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(`{"title":"ship API"}`))
	rec := httptest.NewRecorder()

	CreateTask(store, rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Fatalf("content-type = %q, want application/json", ct)
	}
	var task Task
	if err := json.Unmarshal(rec.Body.Bytes(), &task); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if task.Title != "ship API" {
		t.Fatalf("title = %q, want %q", task.Title, "ship API")
	}
	if task.ID == "" {
		t.Fatal("expected an assigned id")
	}
}

func TestCreateTaskRejectsBadInput(t *testing.T) {
	store := newMemoryStore()
	payloads := []string{
		`{"title":"   "}`,   // blank title
		`{"title":"x","extra":1}`, // unknown field
		`not json`,          // malformed
	}
	for _, payload := range payloads {
		req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(payload))
		rec := httptest.NewRecorder()
		CreateTask(store, rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("payload %q: status = %d, want 400", payload, rec.Code)
		}
	}
}

func TestListTasks(t *testing.T) {
	store := newMemoryStore()
	if _, err := store.Create(Task{Title: "one"}); err != nil {
		t.Fatalf("seed: %v", err)
	}
	if _, err := store.Create(Task{Title: "two"}); err != nil {
		t.Fatalf("seed: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	ListTasks(store, rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	var tasks []Task
	if err := json.Unmarshal(rec.Body.Bytes(), &tasks); err != nil {
		t.Fatalf("decode list: %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("len = %d, want 2", len(tasks))
	}
}