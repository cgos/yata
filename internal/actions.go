package internal

import (
	"os"
	"strings"

	"github.com/cgos/yata/model"
	"github.com/olekukonko/tablewriter"
)

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
func Show(label string, store *FileStore) {
	todolist := store.Read()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Completed", "Date", "Title", "Details"})

	rowByLabel := make(map[string][][]string)
	for _, todo := range todolist {
		row := make([]string, 4)
		var done string

		if todo.Completed == true {
			done = "[x]"
		} else {
			done = "[]"
		}
		row = append(row, done)

		row = append(row, todo.CreatedDate)
		row = append(row, todo.Title)

		details := strings.Join(todo.Details, "\n")
		row = append(row, details)

		for _, label := range todo.Labels {
			rowByLabel[label] = append(rowByLabel[label], row)
		}

		table.Append(row)
	}

	table.Render()
}
