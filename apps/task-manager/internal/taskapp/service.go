package taskapp

import (
	"context"
	"time"

	"github.com/LDKhangg/go-playground/apps/task-manager/internal/tasks"
)

type store interface {
	Add(title string) (tasks.Task, error)
	List() []tasks.Task
	Get(id int) (tasks.Task, error)
	Update(id int, title *string, done *bool) (tasks.Task, error)
	Delete(id int) error
}

type sleeper interface {
	Sleep(ctx context.Context, delay time.Duration) error
}

type realSleeper struct{}

func (realSleeper) Sleep(ctx context.Context, delay time.Duration) error {
	if delay <= 0 {
		return ctx.Err()
	}

	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

type Service struct {
	store   store
	sleeper sleeper
}

func NewService(store *tasks.Store) *Service {
	return &Service{store: store, sleeper: realSleeper{}}
}

func (s *Service) List(ctx context.Context, delay time.Duration) ([]tasks.Task, error) {
	if err := s.sleeper.Sleep(ctx, delay); err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return s.store.List(), nil
}

func (s *Service) Create(ctx context.Context, title string) (tasks.Task, error) {
	if err := ctx.Err(); err != nil {
		return tasks.Task{}, err
	}
	return s.store.Add(title)
}

func (s *Service) Get(ctx context.Context, id int) (tasks.Task, error) {
	if err := ctx.Err(); err != nil {
		return tasks.Task{}, err
	}
	return s.store.Get(id)
}

func (s *Service) Update(ctx context.Context, id int, title *string, done *bool) (tasks.Task, error) {
	if err := ctx.Err(); err != nil {
		return tasks.Task{}, err
	}
	return s.store.Update(id, title, done)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.store.Delete(id)
}
