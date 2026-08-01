package methods

type Task struct {
	ID    int
	Title string
}

func (t *Task) Rename(title string) error {
	panic("TODO")
}

func (t Task) Summary() string {
	panic("TODO")
}
