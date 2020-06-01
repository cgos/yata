package model

import "fmt"

// Todo structure for yata application
type Todo struct {
	Title     string   `json:"title"`
	Completed bool     `json:"completed"`
	Labels    []string `json:"labels"`
	Details   []string `json:"details"`
}

// PrintTodo returns a formatted string
func (t Todo) PrintTodo() string {
	s := fmt.Sprintf("%s \t %t \n Labels: %s \n Details: %s \n", t.Title, t.Completed, t.Labels, t.Details)
	return s
}
