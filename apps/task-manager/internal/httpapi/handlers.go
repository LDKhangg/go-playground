package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/LDKhangg/go-playground/apps/task-manager/internal/taskapp"
	"github.com/LDKhangg/go-playground/apps/task-manager/internal/tasks"
)

type createTaskRequest struct {
	Title string `json:"title"`
}

type patchTaskRequest struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func TasksHandler(service *taskapp.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			delay, err := requestDelay(r)
			if err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid delay"})
				return
			}

			list, err := service.List(r.Context(), delay)
			if err != nil {
				writeTaskError(w, err)
				return
			}
			writeJSON(w, http.StatusOK, list)
		case http.MethodPost:
			var req createTaskRequest
			if err := decodeJSONBody(r, &req); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
				return
			}

			task, err := service.Create(r.Context(), req.Title)
			if err != nil {
				writeTaskError(w, err)
				return
			}

			writeJSON(w, http.StatusCreated, task)
		default:
			w.Header().Set("Allow", "GET, POST")
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	}
}

func TaskByIDHandler(service *taskapp.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := taskIDFromPath(r.URL.Path)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid task id"})
			return
		}

		switch r.Method {
		case http.MethodGet:
			task, err := service.Get(r.Context(), id)
			if err != nil {
				writeTaskError(w, err)
				return
			}
			writeJSON(w, http.StatusOK, task)
		case http.MethodPatch:
			var req patchTaskRequest
			if err := decodeJSONBody(r, &req); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
				return
			}
			if req.Title == nil && req.Done == nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "request must update at least one field"})
				return
			}

			task, err := service.Update(r.Context(), id, req.Title, req.Done)
			if err != nil {
				writeTaskError(w, err)
				return
			}
			writeJSON(w, http.StatusOK, task)
		case http.MethodDelete:
			if err := service.Delete(r.Context(), id); err != nil {
				writeTaskError(w, err)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			w.Header().Set("Allow", "GET, PATCH, DELETE")
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	}
}

func requestDelay(r *http.Request) (time.Duration, error) {
	raw := r.URL.Query().Get("delay")
	if raw == "" {
		return 0, nil
	}

	return time.ParseDuration(raw)
}

func decodeJSONBody(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(dst); err != nil {
		return err
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return errors.New("invalid json")
	}
	return nil
}

func taskIDFromPath(path string) (int, error) {
	idText := strings.TrimPrefix(path, "/tasks/")
	if idText == "" || strings.Contains(idText, "/") {
		return 0, errors.New("invalid task id")
	}

	id, err := strconv.Atoi(idText)
	if err != nil || id < 1 {
		return 0, errors.New("invalid task id")
	}

	return id, nil
}

func writeTaskError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, context.Canceled), errors.Is(err, context.DeadlineExceeded):
		status = http.StatusRequestTimeout
	case errors.Is(err, tasks.ErrEmptyTitle):
		status = http.StatusBadRequest
	case errors.Is(err, tasks.ErrTaskNotFound):
		status = http.StatusNotFound
	}

	writeJSON(w, status, map[string]string{"error": err.Error()})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
