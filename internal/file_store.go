package internal

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/cgos/yata/model"
)

// FileStore used to manage todos in file and app configs
type FileStore struct {
	FilePath string
}

func (f *FileStore) Read() ([]*model.Todo, error) {
	// data, err := ioutil.ReadFile(f.FilePath)
	return nil, nil
}

func (f *FileStore) Write(todolist []*model.Todo) {
	file, err := os.Create(f.FilePath)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	data, err := json.Marshal(todolist)
	if err != nil {
		panic("Unable to marshal todolist")
	}

	nn, err := w.Write(data)
	if (nn < len(data)) || err != nil {
		panic("Error while writing to file")
	}
}
