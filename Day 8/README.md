
# REST API for Task Manager with Authentication

This project is a simple REST API for a task manager written in Go. It allows users to sign up, log in, and manage their tasks through a web interface. The project demonstrates how to implement user authentication with JWT tokens and provides CRUD operations for managing tasks. The API is designed to be simple and straightforward, showcasing the use of Go's `net/http` package for handling requests and responses.

## How It Works

The program works as follows:

1. **User Authentication**:
   - Users can sign up by providing a username and password. The password is securely hashed using bcrypt before being stored in memory.
   - Users can log in by entering their credentials, which are verified against the stored data. Upon successful login, a JWT token is generated and sent back to the user.

2. **Task Management**:
   - Authenticated users can create, read, update, and delete tasks. Each task consists of a title and content.
   - Tasks are stored in memory and are associated with the logged-in user.

3. **Routing and Middleware**:
   - The API uses the Gorilla Mux router for handling different routes. Middleware is implemented to protect routes that require authentication.
   - The API has endpoints for user signup, login, and task management.

4. **Static Files**:
   - The API serves static HTML pages for the frontend. Users can access the signup and login pages through their web browser.

## Project Structure

```
Day 8/
├── go.mod                  # Go module file (already initialized for you)
├── main.go                 # Main file containing API and static file handling
├── handlers/               # Folder containing HTTP handlers for authentication and tasks
│   ├── auth.go             # User authentication logic
│   └── task.go             # Task management logic
├── static/                 # Folder containing HTML and JavaScript files
│   ├── index.html          # Home page with links to signup and login
│   ├── signup.html         # Signup page for new users
│   ├── login.html          # Login page for existing users
│   ├── dashboard.html      # Dashboard for managing tasks
│   ├── main.js             # JavaScript file for handling user interactions
│   └── css/                # Folder containing css styles
└        └── styles.css     # CSS file for styling the HTML pages
```

## How to Run the Program

### 1. Clone the Repository

First, clone the repository and navigate to the `Day 8` directory:

```bash
git clone https://github.com/iovossos/30_Day_Challenge.git
cd 30_Day_Challenge/Day\ 8
```

### 2. Run the Program

The project is already initialized with `go.mod`, so there's no need to initialize it again. To start the API server, run the following command:

```bash
go run main.go
```

### 3. Access the Web Interface

Open your web browser and go to `http://localhost:8000/static/index.html` to access the application.

### Example Usage

- **Sign Up**: Navigate to the signup page, enter a username and password, and click "Sign Up". If the username already exists, an error message will be displayed.
  
- **Login**: After signing up, navigate to the login page, enter your credentials, and click "Login". If successful, you will be redirected to the dashboard.

- **Create Task**: On the dashboard, enter a task title and content, then click "Create Task". The task will be added to your task list.

- **View Tasks**: You can view your tasks directly on the dashboard.

- **Log Out**: Clicking the "Log Out" button will redirect you to the login page.

## File Breakdown

### `main.go`

- **`main()`**: The entry point of the program. It sets up routing, serves static files, and starts the HTTP server.

### `auth.go` (in the `handlers` package)

- **`Signup(w http.ResponseWriter, r *http.Request)`**: Handles user signup, hashes the password, and stores the user in memory.
- **`Login(w http.ResponseWriter, r *http.Request)`**: Handles user login, verifies credentials, and issues a JWT token.
- **`JWTAuthMiddleware(next http.Handler)`**: Middleware that checks for the presence of a valid JWT token for protected routes.

### `task.go` (in the `handlers` package)

- **`CreateTask(w http.ResponseWriter, r *http.Request)`**: Handles creating new tasks.
- **`GetTasks(w http.ResponseWriter, r *http.Request)`**: Retrieves the list of tasks for the authenticated user.
- **`UpdateTask(w http.ResponseWriter, r *http.Request)`**: Updates an existing task.
- **`DeleteTask(w http.ResponseWriter, r *http.Request)`**: Deletes a specified task.

### `index.html`, `signup.html`, `login.html`, `dashboard.html` (in the `static` folder)

- These HTML files provide the user interface for the application, allowing users to sign up, log in, and manage their tasks.

### `main.js`

- Contains JavaScript to handle user interactions, such as submitting forms and fetching tasks.

## Limitations

- **Task Data Persistence**: The tasks are stored in memory, meaning all data will be lost when the server is restarted. For a production application, a database should be implemented.
- **User Data Persistence**: User credentials are stored in memory. Again, for persistence, a database would be necessary.
- **Task Visibility**: Due to lack of current knowledge and time, the implementation does not support having different task lists for different users. All users share the same task storage.
- **Frontend**: The HTML/CSS could be improved for better user experience and design.

## Future Enhancements

1. **Persistent Storage**: Implement a database (e.g., PostgreSQL, MongoDB) to store users and tasks permanently.
2. **User Interface Improvements**: Enhance the HTML and CSS for a better user experience.
3. **Task Sharing**: Allow users to share tasks with each other or have different users see their own tasks.

