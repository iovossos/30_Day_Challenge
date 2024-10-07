package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var tasks []Task
var nextID = 1

// CreateTask handles the POST request to create a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid task data", http.StatusBadRequest)
		return
	}

	task.ID = nextID
	nextID++

	tasks = append(tasks, task)

	log.Printf("Task %d created: %s", task.ID, task.Title)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// ListTasks handles the GET request to list all tasks
func ListTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
