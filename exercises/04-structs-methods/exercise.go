package structsmethods

type Task struct {
	title     string
	completed bool
}

type Project struct {
	current Task
}

func NewTask(title string) Task {
	return Task{}
}

func (t *Task) Complete() {}

func (t Task) IsComplete() bool {
	return false
}

func NewProject(currentTitle string) Project {
	return Project{}
}

func (p *Project) CompleteCurrent() {}

func (p Project) IsCurrentComplete() bool {
	return false
}
