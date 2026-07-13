package structsmethods

type Task struct {
	title     string
	completed bool
}

func NewTask(title string) Task {
	return Task{}
}

func (t *Task) Complete() {}

func (t Task) IsComplete() bool {
	return false
}
