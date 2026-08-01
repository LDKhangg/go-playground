package main

import (
	"net/http"
	"os"
	"time"

	"github.com/LDKhangg/go-playground/apps/task-manager/internal/httpapi"
	"github.com/LDKhangg/go-playground/apps/task-manager/internal/tasks"
)

func serverAddress() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}

	return ":8080"
}

func newMux(store *tasks.Store) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", httpapi.HealthHandler)
	mux.HandleFunc("/tasks", httpapi.TasksHandler(store))
	mux.HandleFunc("/tasks/", httpapi.TaskByIDHandler(store))
	return mux
}

func newServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
}
