package todo

import (
	"encoding/json"
	"os"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func LoadTodos(filename string) ([]Todo, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []Todo{}, nil
	}

	var todos []Todo
	err = json.Unmarshal(data, &todos)
	return todos, err
}

func SaveTodos(filename string, todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func AddTodo(todos []Todo, title string) []Todo {
	id := 1
	if len(todos) > 0 {
		id = todos[len(todos)-1].ID + 1
	}

	newTodo := Todo{
		ID:    id,
		Title: title,
		Done:  false,
	}

	return append(todos, newTodo)
}
