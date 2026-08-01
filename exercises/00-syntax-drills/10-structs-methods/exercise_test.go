//go:build exercise

package methods

import "testing"

func TestTaskRename(t *testing.T) {
	task := Task{ID: 7, Title: "draft"}
	if err := task.Rename("  ship API  "); err != nil {
		t.Fatalf("Rename returned error: %v", err)
	}
	if task.Title != "ship API" {
		t.Fatalf("Title = %q, want %q", task.Title, "ship API")
	}
}

func TestTaskRenameRejectsBlankTitle(t *testing.T) {
	task := Task{ID: 7, Title: "draft"}
	if err := task.Rename("   "); err == nil {
		t.Fatal("expected blank title error")
	}
}

func TestTaskSummary(t *testing.T) {
	task := Task{ID: 7, Title: "ship API"}
	if got := task.Summary(); got == "" {
		t.Fatal("expected non-empty summary")
	}
}
