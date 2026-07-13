package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LDKhangg/go-playground/internal/tasks"
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

	TasksHandler(tasks.NewStore())(recorder, request)

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

	TasksHandler(tasks.NewStore())(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, recorder.Code)
	}
	if got := recorder.Body.String(); got != "{\"id\":1,\"title\":\"learn handlers\"}\n" {
		t.Fatalf("expected created task JSON, got %q", got)
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

			TasksHandler(tasks.NewStore())(recorder, request)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
			}
			if got := recorder.Body.String(); got != tt.want {
				t.Fatalf("expected body %q, got %q", tt.want, got)
			}
		})
	}
}

func TestTasksHandlerRejectsUnsupportedMethod(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/tasks", nil)

	TasksHandler(tasks.NewStore())(recorder, request)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
	if got := recorder.Header().Get("Allow"); got != "GET, POST" {
		t.Fatalf("expected Allow header %q, got %q", "GET, POST", got)
	}
}
