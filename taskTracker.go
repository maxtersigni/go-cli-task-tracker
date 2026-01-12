package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int64
	Description string
	Status      string
}

func AddTask(path string, t Task) error {
	// Read existing file (if it doesn't exist, start empty)
	b, err := os.ReadFile(path)
	var tasks []Task

	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &tasks); err != nil {
			return fmt.Errorf("invalid JSON in %s: %w", path, err)
		}
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Add new task to end
	tasks = append(tasks, t)

	out, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0644)
}

func main() {
	task := Task{
		ID:          2,
		Description: "Goodbye",
		Status:      "Completed",
	}

	if err := AddTask("tasks.json", task); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task added.")
}
