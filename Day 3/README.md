
# Simple Web Server with Custom Routing

This project is a simple web server written in Go. It handles multiple routes and serves static HTML files for each route. The server was built to understand how HTTP routing works in Go without using any external libraries. 

## How It Works

The program consists of multiple route handlers:

1. **`homeHandler`**: This serves the home page and displays a simple HTML file for the root (`/`) route.
2. **`aboutHandler`**: This serves the about page and displays an HTML file for the `/about` route.
3. **`contactHandler`**: This serves the contact page and displays an HTML file for the `/contact` route.
4. **`notFoundHandler`**: This handles any undefined routes and returns a 404 error message.

For each route, the Go server serves the appropriate static HTML file from the `static` folder. The server listens on `localhost:8080`.

## Project Structure

```
Day 3/
├── main.go              # Main file that contains all route handling logic
├── static/              # Folder containing all HTML files
│   ├── index.html       # HTML file for the home page
│   ├── about.html       # HTML file for the about page
│   └── contact.html     # HTML file for the contact page
```

## How to Run the Program

### 1. Clone the Repository
Clone this repository and navigate to the `Day 3` directory.

```bash
git clone https://github.com/iovossos/30_Day_Challenge.git
cd 30_Day_Challenge/Day\ 3
```

### 2. Run the Program
Make sure you have Go installed, and then run the following command:

```bash
go run main.go
```

### Example Usage

You can start the server and visit different routes in your browser.

For example:
- **Home Page**: [http://localhost:8080/](http://localhost:8080/)
- **About Page**: [http://localhost:8080/about](http://localhost:8080/about)
- **Contact Page**: [http://localhost:8080/contact](http://localhost:8080/contact)

The server will display the respective HTML page for each route.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/contact", contactHandler)

    fmt.Println("Server starting on port 8080...")
    http.ListenAndServe(":8080", nil)
}
```

### Output

When visiting each route, you will see the appropriate HTML content.

- Home Page: Displays "Welcome to the Home Page."
- About Page: Displays information about the project.
- Contact Page: Displays contact details.

If you visit an undefined route (e.g., `http://localhost:8080/undefined`), you will see the custom 404 error page:

```
404 - Page Not Found
```

## File Breakdown

### `main.go`

Contains all the logic:

- **`homeHandler(w http.ResponseWriter, r *http.Request)`**: Serves the HTML file for the home page.
- **`aboutHandler(w http.ResponseWriter, r *http.Request)`**: Serves the HTML file for the about page.
- **`contactHandler(w http.ResponseWriter, r *http.Request)`**: Serves the HTML file for the contact page.
- **`notFoundHandler(w http.ResponseWriter, r *http.Request)`**: Serves a 404 error for any undefined routes.
