# Todo List Manager

A command-line to-do list manager built in Go that allows users to add, remove, complete, and list tasks. The tasks are stored in a JSON file, allowing them to persist between sessions.

## Features
- Add new tasks to the to-do list
- List all tasks with completion status
- Mark tasks as complete
- Remove tasks from the list
- Persistent storage using a JSON file

## Installation

### Prerequisites
- You must have [Go](https://golang.org/dl/) installed on your machine.

### Steps
1. Clone the repository:
    ```bash
    git clone <repository-url>
    cd todo-list
    ```

2. Build the program:
    ```bash
    go build -o todo
    ```

3. You now have a `todo` executable that you can run from the terminal.

## Usage

### Add a Task
To add a task, use the following command:
```bash
./todo add "Your task description"
