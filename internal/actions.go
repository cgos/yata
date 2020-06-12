package internal

import "github.com/cgos/yata/model"

// Add actions read/update the yata file with a new item
func Add(item *model.Todo, store *FileStore) {
	if item == nil {
		return
	}
	todolist := store.Read()
	todolist = append(todolist, item)
	store.Write(todolist)
}

// Show action list all todos related to the label. If no labels, then all todos are returned
func Show(label string, store *FileStore) []*model.Todo {
	todolist := store.Read()
	return todolist
}
