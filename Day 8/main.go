package main

import (
	"log"
	"net/http"
	"taskmanager/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Serve static files (HTML, CSS, JS) from the /static/ folder
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Public routes for signup and login
	r.HandleFunc("/signup", handlers.Signup).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected route for the dashboard (only accessible with valid JWT)
	r.Handle("/static/dashboard.html", handlers.JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/dashboard.html")
	})))

	// Task management routes (protected with JWT) located in handlers/task.go
	r.Handle("/api/tasks", handlers.JWTAuthMiddleware(http.HandlerFunc(handlers.CreateTask))).Methods("POST")
	r.Handle("/api/tasks", handlers.JWTAuthMiddleware(http.HandlerFunc(handlers.ListTasks))).Methods("GET")

	// Start the server on port 8000
	log.Println("Server starting on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
