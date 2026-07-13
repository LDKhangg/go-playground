package httpapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/LDKhangg/go-playground/internal/tasks"
)

type createTaskRequest struct {
	Title string `json:"title"`
}

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func TasksHandler(store *tasks.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, store.List())
		case http.MethodPost:
			var req createTaskRequest
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&req); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
				return
			}
			if err := decoder.Decode(&struct{}{}); err != io.EOF {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
				return
			}

			task, err := store.Add(req.Title)
			if err != nil {
				status := http.StatusInternalServerError
				if errors.Is(err, tasks.ErrEmptyTitle) {
					status = http.StatusBadRequest
				}
				writeJSON(w, status, map[string]string{"error": err.Error()})
				return
			}

			writeJSON(w, http.StatusCreated, task)
		default:
			w.Header().Set("Allow", "GET, POST")
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
