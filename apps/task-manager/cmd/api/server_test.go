package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/LDKhangg/go-playground/apps/task-manager/internal/taskapp"
	"github.com/LDKhangg/go-playground/apps/task-manager/internal/tasks"
)

func TestServerAddressDefaultsTo8080(t *testing.T) {
	t.Setenv("PORT", "")

	if got := serverAddress(); got != ":8080" {
		t.Fatalf("serverAddress() = %q, want %q", got, ":8080")
	}
}

func TestServerAddressUsesPortEnv(t *testing.T) {
	t.Setenv("PORT", "9090")

	if got := serverAddress(); got != ":9090" {
		t.Fatalf("serverAddress() = %q, want %q", got, ":9090")
	}
}

func TestNewMuxRegistersCollectionAndItemRoutes(t *testing.T) {
	store := tasks.NewStore()
	if _, err := store.Add("learn shutdown"); err != nil {
		t.Fatalf("Add returned error: %v", err)
	}

	handler := newMux(taskapp.NewService(store))

	listRequest := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	listRecorder := httptest.NewRecorder()
	handler.ServeHTTP(listRecorder, listRequest)
	if listRecorder.Code != http.StatusOK {
		t.Fatalf("GET /tasks status = %d, want %d", listRecorder.Code, http.StatusOK)
	}

	itemRequest := httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	itemRecorder := httptest.NewRecorder()
	handler.ServeHTTP(itemRecorder, itemRequest)
	if itemRecorder.Code != http.StatusOK {
		t.Fatalf("GET /tasks/1 status = %d, want %d", itemRecorder.Code, http.StatusOK)
	}
}

func TestNewServerConfiguresTimeouts(t *testing.T) {
	server := newServer(":9090", http.NewServeMux())

	if server.Addr != ":9090" {
		t.Fatalf("Addr = %q, want %q", server.Addr, ":9090")
	}
	if server.ReadHeaderTimeout != 5*time.Second {
		t.Fatalf("ReadHeaderTimeout = %s, want %s", server.ReadHeaderTimeout, 5*time.Second)
	}
	if server.IdleTimeout != 30*time.Second {
		t.Fatalf("IdleTimeout = %s, want %s", server.IdleTimeout, 30*time.Second)
	}
}
