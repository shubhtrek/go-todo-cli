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
