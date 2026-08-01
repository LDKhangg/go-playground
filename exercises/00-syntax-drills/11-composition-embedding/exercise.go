package embedding

type Task struct {
	Title string
}

func (t Task) Summary() string {
	return t.Title
}

type TimedTask struct {
	Task
	DueInHours int
}

func (t TimedTask) Label() string {
	panic("TODO")
}
