package todosRepository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/S3ergio31/task-cli/todos"
)

func LoadTodos(todos *map[string]todos.Task) {
	jsonFile, err := os.Open("todos.json")

	if err != nil {
		fmt.Printf("loadTodos: %s\n", err)

		return
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &todos)
}

func SaveTodos(todos *map[string]todos.Task) {
	bytes, jsonErr := json.Marshal(todos)

	if jsonErr != nil {
		fmt.Printf("saveTodos: %s\n", jsonErr)

		return
	}

	writeError := os.WriteFile("todos.json", bytes, 0644)

	if writeError != nil {
		fmt.Printf("saveTodos: %s\n", writeError)

		return
	}
}
