package main

import (
	"os"
	"path/filepath"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	dataFilePath := filepath.Join(homeDir, ".gotodo", "todos.json")
	os.MkdirAll(filepath.Dir(dataFilePath), os.ModePerm)

	todos := Todos{}

	storage := NewStorage[Todos](dataFilePath)
	storage.Load(&todos)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	storage.Save(todos)
}
