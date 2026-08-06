package jsonbasics

import (
	"encoding/json"
	"errors"
	"strings"
)

type Task struct {
	Title string `json:"title"`
}

func DecodeTask(data []byte) (Task, error) {
	var m Task
	err := json.Unmarshal(data, &m)
	if err != nil {
		return Task{}, err
	}
	trimmedTitle := strings.TrimSpace(m.Title)
	if trimmedTitle == "" {
		return Task{}, errors.New("Invalid title")
	}
	return m, nil
}
