package internal

import "github.com/cgos/yata/model"

// FileStore used to manage todos in file
type FileStore struct {
	FilePath string
}

func (f *FileStore) Read() ([]*model.Todo, error) {
	return nil, nil
}

func (f *FileStore) Write(todolist []*model.Todo) {

}
