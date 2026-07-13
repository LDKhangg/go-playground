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
