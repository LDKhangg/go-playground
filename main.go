package main

import (
	"log"
	"net/http"

	"github.com/LDKhangg/go-playground/internal/httpapi"
	"github.com/LDKhangg/go-playground/internal/tasks"
)

func main() {
	store := tasks.NewStore()
	store.Add("learn go")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", httpapi.HealthHandler)
	mux.HandleFunc("/tasks", httpapi.TasksHandler(store))

	log.Println("listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
