package internal

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cgos/yata/model"
)

// FileStore used to manage todos in file and app configs
type FileStore struct {
	yataFilePath string
}

// NewFileStore creates the FileStore struct
func NewFileStore(path string) *FileStore {
	if len(path) == 0 {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			panic("Unable to get home dir")
		}
		path = userHomeDir
	}

	path = filepath.Join(path, ".yata.json")
	return &FileStore{yataFilePath: path}
}

func (f *FileStore) Read() []*model.Todo {
	data, err := ioutil.ReadFile(f.yataFilePath)
	todolist := []*model.Todo{}
	if err != nil {
		return todolist
	}

	err = json.Unmarshal(data, &todolist)
	if err != nil {
		panic("Unable to unmarshal json from file")
	}

	return todolist
}

func (f *FileStore) Write(todolist []*model.Todo) {
	if len(todolist) == 0 {
		return
	}

	file, err := os.Create(f.yataFilePath)
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
