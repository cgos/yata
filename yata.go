package main

import (
	"fmt"

	"github.com/cgos/yata/model"
)

func main() {
	fmt.Printf("hello toto")
	t := model.Todo{Title: "allo"}
	fmt.Printf("%s", t.Title)
}
