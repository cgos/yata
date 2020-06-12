package internal

import (
	"os"
	"testing"

	"github.com/cgos/yata/model"
)

func TestAdd(t *testing.T) {
	fs := &FileStore{yataFilePath: "test-yata.json"}
	defer os.Remove(fs.yataFilePath)

	todo := &model.Todo{Title: "new-todo", Completed: false, Labels: []string{"Main"}, Details: []string{"This section is meant to include some info about the task"}}

	Add(todo, fs)
	todos := fs.Read()
	if len(todos) != 1 {
		t.Error("list of todos written not same size as read")
	}
}

func TestAddEmpty(t *testing.T) {
	fs := &FileStore{yataFilePath: "test-yata.json"}
	defer os.Remove(fs.yataFilePath)

	Add(nil, fs)

	todos := fs.Read()
	if len(todos) != 0 {
		t.Error("list of todos written not same size as read")
	}
}

complete tests