package utils

import (
	"encoding/json"
	"errors"
	"os"
	"todolist_cli/model"
)

const TodosFilePath = "data/todos.json"

func EnsureTodosFile() error {
	_, err := os.Stat(TodosFilePath)
	if errors.Is(err, os.ErrNotExist) {
		// create folder kalau blm ada
		if err := os.MkdirAll("data", 0755); err != nil {
			return err
		}
		// buat file todos.json dlm array "[]"
		return os.WriteFile(TodosFilePath, []byte("[]"), 0644)
	}
	return nil
}

func LoadTodoFromFile() ([]model.Todos, error) {
	// memastikan file ada
	if err := EnsureTodosFile(); err != nil {
		return nil, err
	}

	file, err := os.ReadFile(TodosFilePath)
	if err != nil {
		return nil, err
	}
	if len(file) == 0 {
		return []model.Todos{}, nil
	}

	var tickets []model.Todos
	if err := json.Unmarshal(file, &tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

// save todo dalam slice []
func SavedTodoToFile(todos []model.Todos) error {
	if err := EnsureTodosFile(); err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(TodosFilePath, bytes, 0644)
}
