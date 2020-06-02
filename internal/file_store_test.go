package internal

import (
	"os"
	"testing"

	"github.com/cgos/yata/model"
)

var todolist []*model.Todo

func BuildTodoFixture() []*model.Todo {
	todo1 := &model.Todo{Title: "first-todo", Completed: false, Labels: []string{"Main"}, Details: []string{"This section is meant to include some info about the task"}}
	todo2 := &model.Todo{Title: "second-todo", Completed: false, Labels: []string{"Main"}, Details: []string{"Some info"}}
	todo3 := &model.Todo{Title: "third-todo", Completed: false, Labels: []string{"Main"}, Details: []string{"More details for this task"}}
	todo4 := &model.Todo{Title: "fourth-todo", Completed: false, Labels: []string{"RandomList", "Main"}, Details: []string{"Some random info here", "a second notes on this tasks"}}

	todos := []*model.Todo{todo1, todo2, todo3, todo4}
	return todos
}

func TestRead(t *testing.T) {
	fs := FileStore{FilePath: "test-yata.json"}
	fs.Write(todolist)
	defer os.Remove(fs.FilePath)

	todos := fs.Read()
	if len(todos) != len(todolist) {
		t.Errorf("list of todos written not same size as read")
	}
}

func TestWrite(t *testing.T) {
	fs := FileStore{FilePath: "test-yata.json"}
	fs.Write(todolist)
	defer os.Remove(fs.FilePath)
	for _, item := range todolist {
		s := item.PrintTodo()
		t.Logf("%s", s)
	}
}

func TestMain(m *testing.M) {
	//check how to use subtest: https://www.reddit.com/r/golang/comments/axerii/using_testmain_on_multiple_series_of_tests/
	todolist = BuildTodoFixture()
	m.Run()
	os.Exit(1)
}
