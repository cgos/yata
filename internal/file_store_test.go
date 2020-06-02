package internal

import (
	"testing"

	"github.com/cgos/yata/model"
)

func BuildTodoFixture() []*model.Todo {
	todo1 := &model.Todo{Title: "first-todo", Completed: false, Labels: []string{"Main"}, Details: []string{"This section is meant to include some info about the task"}}
	todo2 := &model.Todo{Title: "second-todo", Completed: false, Labels: []string{"Main"}, Details: []string{"Some info"}}
	todo3 := &model.Todo{Title: "third-todo", Completed: false, Labels: []string{"Main"}, Details: []string{"More details for this task"}}
	todo4 := &model.Todo{Title: "fourth-todo", Completed: false, Labels: []string{"RandomList", "Main"}, Details: []string{"Some random info here", "a second notes on this tasks"}}

	todolist := []*model.Todo{todo1, todo2, todo3, todo4}
	return todolist
}

func TestWrite(t *testing.T) {
	// store := &FileStore{FilePath: "yata-test.json"}
	todolist := BuildTodoFixture()
	fs := FileStore{FilePath: "test.yata"}
	fs.Write(todolist)
	for _, item := range todolist {
		s := item.PrintTodo()
		t.Logf("%s", s)
	}
}
