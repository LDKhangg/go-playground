//go:build exercise

package embedding

import (
	"strings"
	"testing"
)

func TestTimedTaskLabel(t *testing.T) {
	task := TimedTask{Task: Task{Title: "ship API"}, DueInHours: 4}
	label := task.Label()
	if !strings.Contains(label, "ship API") {
		t.Fatalf("label = %q, want embedded title", label)
	}
	if !strings.Contains(label, "4") {
		t.Fatalf("label = %q, want due hours", label)
	}
}
