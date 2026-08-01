package oopdesign

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrEmptyTitle   = errors.New("title must not be empty")
	ErrTaskNotFound = errors.New("task not found")
)

type Task struct {
	ID          string
	Title       string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

func (t *Task) MarkDone(now time.Time) {
	panic("TODO")
}

type Repository interface {
	Save(task Task) error
	FindByID(id string) (Task, error)
}

type Clock interface {
	Now() time.Time
}

type IDGenerator interface {
	NextID() string
}

type TaskService struct {
	repo  Repository
	clock Clock
	ids   IDGenerator
}

func NewTaskService(repo Repository, clock Clock, ids IDGenerator) *TaskService {
	return &TaskService{repo: repo, clock: clock, ids: ids}
}

func (s *TaskService) Create(title string) (Task, error) {
	_ = strings.TrimSpace(title)
	panic("TODO")
}

func (s *TaskService) Complete(id string) (Task, error) {
	_ = id
	panic("TODO")
}
