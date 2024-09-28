
# Day 1: Todo List Manager

This is the Day 1 project of the 30-Day Go Programming Challenge. It's a simple **command-line Todo List Manager** that allows you to add, list, complete, and remove tasks. The tasks are stored in a JSON file for persistence.

## Features

- Add tasks to your todo list.
- List all tasks, with a marker showing whether the task is completed or not.
- Mark a task as complete.
- Remove tasks from the list.
- All tasks are stored in a JSON file, making them persistent across runs.

## How to Run the Program

### 1. Clone the Repository
Clone this repository and navigate to the `Day 1` directory.

```bash
git clone https://github.com/your-username/30_Day_Challenge.git
cd 30_Day_Challenge/Day\ 1
```

### 2. Run the Program
Make sure you have Go installed, and then run the following command:

```bash
go run Main/Main.go <command> [arguments]
```

### Available Commands

| Command                              | Description                                          |
|--------------------------------------|------------------------------------------------------|
| `go run Main/Main.go add <task>`     | Adds a task to your todo list.                       |
| `go run Main/Main.go list`           | Lists all tasks in your todo list.                   |
| `go run Main/Main.go complete <id>`  | Marks a specific task as completed.                  |
| `go run Main/Main.go remove <id>`    | Removes a task from your todo list.                  |

### Example Usage

1. **Add a task**:
   ```bash
   go run Main/Main.go add "Buy groceries"
   ```

2. **List all tasks**:
   ```bash
   go run Main/Main.go list
   ```

3. **Complete a task**:
   ```bash
   go run Main/Main.go complete 1
   ```

4. **Remove a task**:
   ```bash
   go run Main/Main.go remove 1
   ```

## File Storage

Tasks are stored in a `tasks.json` file in the same directory as the program. This file is created automatically when you first add a task, and it is updated each time you add, remove, or mark tasks as complete.

### Example `tasks.json` file

```json
{
  "tasks": [
    {
      "description": "Buy groceries",
      "completed": false
    },
    {
      "description": "Walk the dog",
      "completed": true
    }
  ]
}
```

## Requirements

- **Go**: You need Go installed to run this program. You can download it [here](https://golang.org/dl/).

## License

This project is open source under the [MIT License](LICENSE).
