//go:build exercise

package gormlearn

import (
	"errors"
	"fmt"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	name := fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	if err := db.AutoMigrate(&Task{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func TestCreateAndFindTask(t *testing.T) {
	db := newTestDB(t)

	task, err := CreateTask(db, "learn gorm")
	if err != nil {
		t.Fatalf("CreateTask: %v", err)
	}
	if task.Title != "learn gorm" {
		t.Fatalf("title = %q, want %q", task.Title, "learn gorm")
	}
	if task.ID == 0 {
		t.Fatal("expected an auto-assigned id")
	}

	got, err := FindTask(db, task.ID)
	if err != nil {
		t.Fatalf("FindTask: %v", err)
	}
	if got.Title != task.Title {
		t.Fatalf("found title = %q, want %q", got.Title, task.Title)
	}
}

func TestFindMissingReturnsErrRecordNotFound(t *testing.T) {
	db := newTestDB(t)

	if _, err := FindTask(db, 999); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("err = %v, want gorm.ErrRecordNotFound", err)
	}
}

func TestMarkDone(t *testing.T) {
	db := newTestDB(t)

	task, err := CreateTask(db, "ship api")
	if err != nil {
		t.Fatalf("CreateTask: %v", err)
	}

	done, err := MarkDone(db, task.ID)
	if err != nil {
		t.Fatalf("MarkDone: %v", err)
	}
	if !done.Done {
		t.Fatal("expected done=true")
	}

	again, err := FindTask(db, task.ID)
	if err != nil {
		t.Fatalf("FindTask: %v", err)
	}
	if !again.Done {
		t.Fatal("expected done to persist in the database")
	}
}

func TestListTasks(t *testing.T) {
	db := newTestDB(t)

	for _, title := range []string{"one", "two"} {
		if _, err := CreateTask(db, title); err != nil {
			t.Fatalf("seed %q: %v", title, err)
		}
	}

	tasks, err := ListTasks(db)
	if err != nil {
		t.Fatalf("ListTasks: %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("len = %d, want 2", len(tasks))
	}
}