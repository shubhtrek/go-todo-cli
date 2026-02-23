package main

import (
	"bufio"
	"fmt"
	"go-todo-cli/todo"
	"os"
	"strconv"
	"strings"
)

const dataFile = "data.json"

func main() {
	reader := bufio.NewReader(os.Stdin)

	todos, err := todo.LoadTodos(dataFile)
	if err != nil {
		fmt.Println("❌ Error loading todos:", err)
		return
	}

	for {
		fmt.Println("\n📝 Todo CLI App")
		fmt.Println("1. Add Todo")
		fmt.Println("2. List Todos")
		fmt.Println("3. Mark Done")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Exit")
		fmt.Print("Choose option: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			fmt.Print("Enter todo title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			todos = todo.AddTodo(todos, title)
			_ = todo.SaveTodos(dataFile, todos)
			fmt.Println("✅ Todo added")
	}
}
