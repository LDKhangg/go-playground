package jsonbasics

type Task struct {
	Title string `json:"title"`
}

func DecodeTask(data []byte) (Task, error) {
	panic("TODO")
}
