package main

import (
	"fmt"
	"os"
	"strconv"

	"TodoList/task" // Importing the task package
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <command> [arguments]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . add <task description>")
			return
		}
		description := os.Args[2]
		task.AddTask(description)
	case "list":
		task.ListTasks()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . complete <task number>")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task number:", os.Args[2])
			return
		}
		task.CompleteTask(index - 1)
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . remove <task number>")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task number:", os.Args[2])
			return
		}
		task.RemoveTask(index - 1)
	case "help":
		fmt.Println("add        Usage: go run . add <task description>. Used to add tasks on the todo list.")
		fmt.Println("list       Usage: go run . list. Used to list all tasks.")
		fmt.Println("complete   Usage: go run . complete <task number>. Used to mark a task as completed.")
		fmt.Println("remove     Usage: go run . remove <task number>. Used to remove a task from the todo list.")
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Use go run . help to get a list of available commands.")
	}
}
