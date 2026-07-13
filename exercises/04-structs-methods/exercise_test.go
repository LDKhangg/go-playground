//go:build exercise

package structsmethods

import "testing"

func TestTaskLifecycle(t *testing.T) {
	task := NewTask("learn methods")
	if task.title != "learn methods" {
		t.Fatalf("NewTask title = %q, want %q", task.title, "learn methods")
	}
	if task.IsComplete() {
		t.Fatal("new task is complete, want incomplete")
	}

	task.Complete()
	if !task.IsComplete() {
		t.Fatal("task remains incomplete after Complete")
	}
}

func TestProjectComposesAndDelegatesToTask(t *testing.T) {
	project := NewProject("learn composition")
	if project.current.title != "learn composition" {
		t.Fatalf("NewProject current task title = %q, want %q", project.current.title, "learn composition")
	}
	if project.IsCurrentComplete() {
		t.Fatal("new project's current task is complete, want incomplete")
	}

	project.CompleteCurrent()
	if !project.IsCurrentComplete() {
		t.Fatal("current task remains incomplete after CompleteCurrent")
	}
}
