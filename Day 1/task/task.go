package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

// Load tasks from the JSON file using os.ReadFile
func LoadTasks() (TaskList, error) {
	var tasks TaskList
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return TaskList{Tasks: []Task{}}, nil
		}
		return tasks, err
	}
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

// Save tasks to the JSON file using os.WriteFile
func SaveTasks(tasks TaskList) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

// Add a new task to the list
func AddTask(description string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	newTask := Task{
		Description: description,
		Completed:   false,
	}
	tasks.Tasks = append(tasks.Tasks, newTask)

	if err := SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Println("Task added:", description)
}

// List all tasks
func ListTasks() {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for i, task := range tasks.Tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, i+1, task.Description)
	}
}

// Mark a task as completed
func CompleteTask(index int) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if index < 0 || index >= len(tasks.Tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks.Tasks[index].Completed = true

	if err := SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task marked as complete:", tasks.Tasks[index].Description)
}

// Remove a task from the list
func RemoveTask(index int) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if index < 0 || index >= len(tasks.Tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks.Tasks = append(tasks.Tasks[:index], tasks.Tasks[index+1:]...)

	if err := SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task removed.")
}
