package httpapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LDKhangg/go-playground/apps/task-manager/internal/taskapp"
	"github.com/LDKhangg/go-playground/apps/task-manager/internal/tasks"
)

func TestHealthHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/health", nil)

	HealthHandler(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected JSON content type, got %q", got)
	}
	if got := recorder.Body.String(); got != "{\"status\":\"ok\"}\n" {
		t.Fatalf("expected health JSON, got %q", got)
	}
}

func TestTasksHandlerListsEmptyStoreAsArray(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/tasks", nil)

	TasksHandler(taskapp.NewService(tasks.NewStore()))(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
	if got := recorder.Body.String(); got != "[]\n" {
		t.Fatalf("expected empty JSON array, got %q", got)
	}
}

func TestTasksHandlerCreatesTask(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"  learn handlers  "}`))

	TasksHandler(taskapp.NewService(tasks.NewStore()))(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, recorder.Code)
	}
	if got := recorder.Body.String(); got != "{\"id\":1,\"title\":\"learn handlers\",\"done\":false}\n" {
		t.Fatalf("expected created task JSON, got %q", got)
	}
}

func TestTaskByIDHandlerGetsTask(t *testing.T) {
	store := tasks.NewStore()
	service := taskapp.NewService(store)
	if _, err := store.Add("learn handlers"); err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)

	TaskByIDHandler(service)(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
	if got := recorder.Body.String(); got != "{\"id\":1,\"title\":\"learn handlers\",\"done\":false}\n" {
		t.Fatalf("expected task JSON, got %q", got)
	}
}

func TestTaskByIDHandlerPatchesTask(t *testing.T) {
	store := tasks.NewStore()
	service := taskapp.NewService(store)
	if _, err := store.Add("learn handlers"); err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPatch, "/tasks/1", strings.NewReader(`{"title":"ship handlers","done":true}`))

	TaskByIDHandler(service)(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
	if got := recorder.Body.String(); got != "{\"id\":1,\"title\":\"ship handlers\",\"done\":true}\n" {
		t.Fatalf("expected patched task JSON, got %q", got)
	}
}

func TestTaskByIDHandlerDeletesTask(t *testing.T) {
	store := tasks.NewStore()
	service := taskapp.NewService(store)
	if _, err := store.Add("learn handlers"); err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)

	TaskByIDHandler(service)(recorder, request)

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, recorder.Code)
	}
	if got := recorder.Body.String(); got != "" {
		t.Fatalf("expected empty body, got %q", got)
	}
}

func TestTasksHandlerRejectsInvalidRequests(t *testing.T) {
	tests := []struct {
		name string
		body string
		want string
	}{
		{name: "invalid JSON", body: `{`, want: "{\"error\":\"invalid json\"}\n"},
		{name: "trailing malformed JSON", body: `{"title":"learn handlers"}{`, want: "{\"error\":\"invalid json\"}\n"},
		{name: "multiple JSON values", body: `{"title":"learn handlers"}{"title":"ignored"}`, want: "{\"error\":\"invalid json\"}\n"},
		{name: "empty title", body: `{"title":"   "}`, want: "{\"error\":\"title must not be empty\"}\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(tt.body))

			TasksHandler(taskapp.NewService(tasks.NewStore()))(recorder, request)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
			}
			if got := recorder.Body.String(); got != tt.want {
				t.Fatalf("expected body %q, got %q", tt.want, got)
			}
		})
	}
}

func TestTaskByIDHandlerRejectsInvalidRequests(t *testing.T) {
	store := tasks.NewStore()
	service := taskapp.NewService(store)
	if _, err := store.Add("learn handlers"); err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	tests := []struct {
		name   string
		method string
		path   string
		body   string
		status int
		want   string
	}{
		{name: "invalid ID", method: http.MethodGet, path: "/tasks/nope", status: http.StatusBadRequest, want: "{\"error\":\"invalid task id\"}\n"},
		{name: "not found", method: http.MethodGet, path: "/tasks/99", status: http.StatusNotFound, want: "{\"error\":\"task not found\"}\n"},
		{name: "invalid json", method: http.MethodPatch, path: "/tasks/1", body: `{`, status: http.StatusBadRequest, want: "{\"error\":\"invalid json\"}\n"},
		{name: "multiple JSON values", method: http.MethodPatch, path: "/tasks/1", body: `{"done":true}{}`, status: http.StatusBadRequest, want: "{\"error\":\"invalid json\"}\n"},
		{name: "empty patch", method: http.MethodPatch, path: "/tasks/1", body: `{}`, status: http.StatusBadRequest, want: "{\"error\":\"request must update at least one field\"}\n"},
		{name: "empty title", method: http.MethodPatch, path: "/tasks/1", body: `{"title":"   "}`, status: http.StatusBadRequest, want: "{\"error\":\"title must not be empty\"}\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))

			TaskByIDHandler(service)(recorder, request)

			if recorder.Code != tt.status {
				t.Fatalf("expected status %d, got %d", tt.status, recorder.Code)
			}
			if got := recorder.Body.String(); got != tt.want {
				t.Fatalf("expected body %q, got %q", tt.want, got)
			}
		})
	}
}

func TestTasksHandlerRejectsInvalidDelayQuery(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/tasks?delay=soon", nil)

	TasksHandler(taskapp.NewService(tasks.NewStore()))(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}
	if got := recorder.Body.String(); got != "{\"error\":\"invalid delay\"}\n" {
		t.Fatalf("expected invalid delay error, got %q", got)
	}
}

func TestTasksHandlerReturnsRequestTimeoutWhenContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/tasks?delay=50ms", nil).WithContext(ctx)

	TasksHandler(taskapp.NewService(tasks.NewStore()))(recorder, request)

	if recorder.Code != http.StatusRequestTimeout {
		t.Fatalf("expected status %d, got %d", http.StatusRequestTimeout, recorder.Code)
	}
	if got := recorder.Body.String(); got != "{\"error\":\"context canceled\"}\n" {
		t.Fatalf("expected context canceled error, got %q", got)
	}
}

func TestTasksHandlerRejectsUnsupportedMethod(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/tasks", nil)

	TasksHandler(taskapp.NewService(tasks.NewStore()))(recorder, request)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
	if got := recorder.Header().Get("Allow"); got != "GET, POST" {
		t.Fatalf("expected Allow header %q, got %q", "GET, POST", got)
	}
}

func TestTaskByIDHandlerRejectsUnsupportedMethod(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/tasks/1", nil)

	TaskByIDHandler(taskapp.NewService(tasks.NewStore()))(recorder, request)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
	if got := recorder.Header().Get("Allow"); got != "GET, PATCH, DELETE" {
		t.Fatalf("expected Allow header %q, got %q", "GET, PATCH, DELETE", got)
	}
}
