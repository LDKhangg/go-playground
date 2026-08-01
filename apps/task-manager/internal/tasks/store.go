package tasks

import (
	"errors"
	"strings"
	"sync"
)

var ErrEmptyTitle = errors.New("title must not be empty")
var ErrTaskNotFound = errors.New("task not found")

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
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

func (s *Store) Get(id int) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return Task{}, ErrTaskNotFound
}

func (s *Store) Update(id int, title *string, done *bool) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.tasks {
		if s.tasks[i].ID != id {
			continue
		}

		if title != nil {
			trimmed := strings.TrimSpace(*title)
			if trimmed == "" {
				return Task{}, ErrEmptyTitle
			}
			s.tasks[i].Title = trimmed
		}
		if done != nil {
			s.tasks[i].Done = *done
		}

		return s.tasks[i], nil
	}

	return Task{}, ErrTaskNotFound
}

func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.tasks {
		if s.tasks[i].ID != id {
			continue
		}

		s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
		return nil
	}

	return ErrTaskNotFound
}
