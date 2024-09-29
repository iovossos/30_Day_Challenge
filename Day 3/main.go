package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "404 - Page Not Found")
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Define routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	// 404 handler for undefined routes
	http.HandleFunc("/404", notFoundHandler)

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
