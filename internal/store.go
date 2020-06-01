package internal

import "github.com/cgos/yata/model"

// Store interface for Todo
type Store interface {
	Read() ([]*model.Todo, error)
	Write(todolist []*model.Todo)
}
