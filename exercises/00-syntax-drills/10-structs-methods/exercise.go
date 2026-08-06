package methods

import (
	"errors"
	"fmt"
	"strings"
)

type Task struct {
	ID    int
	Title string
}

func (t *Task) Rename(title string) error {
	trimmedStr := strings.TrimSpace(title)

	if trimmedStr == "" {
		return errors.New("title not valid")
	}
	t.Title = trimmedStr
	return nil
}

func (t Task) Summary() string {
	return fmt.Sprintf("%d: %s", t.ID, t.Title)
}
