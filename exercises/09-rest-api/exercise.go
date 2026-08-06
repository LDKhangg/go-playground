package restapi

import (
	"errors"
	"net/http"
)

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Store interface {
	Create(task Task) (Task, error)
	List() ([]Task, error)
	FindByID(id string) (Task, error)
}

var (
	ErrNotFound   = errors.New("task not found")
	ErrEmptyTitle = errors.New("title must not be empty")
)

func CreateTask(store Store, w http.ResponseWriter, r *http.Request) {
	panic("TODO")
}

func ListTasks(store Store, w http.ResponseWriter, r *http.Request) {
	panic("TODO")
}