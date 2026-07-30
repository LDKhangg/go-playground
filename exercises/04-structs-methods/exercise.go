package structsmethods

type Task struct {
	title     string
	completed bool
}

type Project struct {
	current Task
}

func NewTask(title string) Task {
	return Task{
		title: title,
	}
}

func (t *Task) Complete() {
	t.completed = true
}

func (t Task) IsComplete() bool {
	return t.completed
}

func NewProject(currentTitle string) Project {
	return Project{
		current: NewTask(currentTitle),
	}
}

func (p *Project) CompleteCurrent() {
	p.current.Complete()
}

func (p Project) IsCurrentComplete() bool {
	return p.current.IsComplete()
}
