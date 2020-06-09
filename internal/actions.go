package internal

import "github.com/cgos/yata/model"

// Add actions read/update the yata file with a new item
func Add(item *model.Todo) {
	fs := NewFileStore("")
	todolist := fs.Read()
	append(todolist, item)
}
