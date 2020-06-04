package main

import (
	"os"

	"github.com/cgos/yata/cmd"
	"github.com/cgos/yata/internal"
	"github.com/cgos/yata/model"
)

func main() {
	cmd.Execute()
}

// Yata struct to tie everything together
type Yata struct {
	Store *internal.FileStore
	Todos []*model.Todo
}

// NewYata populates the Yata struct
func NewYata() *Yata {
	filepath, err := os.UserHomeDir()
	if err != nil {
		panic("Unable to load UserHomeDir")
	}

	filepath = filepath + os.PathSeparator() + ".yata.json"
	yata := &Yata{
		Store: &internal.FileStore{FilePath: filepath},
		Todos: []*model.Todo{},
	}
	return yata
}
