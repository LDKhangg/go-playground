package tasks

import (
	"errors"
	"strings"
	"sync"
)

var ErrEmptyTitle = errors.New("title must not be empty")

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Store struct {
	mu     sync.Mutex
	nextID int
	tasks  []Task
}

func NewStore() *Store {
	return &Store{
		nextID: 1,
		tasks:  make([]Task, 0),
	}
}

func (s *Store) Add(title string) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, ErrEmptyTitle
	}

	task := Task{ID: s.nextID, Title: title}
	s.nextID++
	s.tasks = append(s.tasks, task)
	return task, nil
}

func (s *Store) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]Task, len(s.tasks))
	copy(out, s.tasks)
	return out
}
