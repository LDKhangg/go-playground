package gormlearn

import "gorm.io/gorm"

var ErrNotFound = gorm.ErrRecordNotFound

type Task struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"size:200;not null"`
	Done  bool
}

func CreateTask(db *gorm.DB, title string) (Task, error) {
	panic("TODO")
}

func FindTask(db *gorm.DB, id uint) (Task, error) {
	panic("TODO")
}

func ListTasks(db *gorm.DB) ([]Task, error) {
	panic("TODO")
}

func MarkDone(db *gorm.DB, id uint) (Task, error) {
	panic("TODO")
}