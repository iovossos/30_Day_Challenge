# Expense Tracker with File Persistence and Reporting

This project is an Expense Tracker application written in Go. It allows users to track their expenses, store them in a JSON file, and generate reports summarizing their spending. The application features a command-line menu interface, providing users with options to add, view, and report on their expenses.

## How It Works

The application is designed to:

1. **Menu-Driven Interface**:
   - The user interacts with the program through a command-line menu. Options include adding new expenses, viewing all expenses, generating reports, and exiting the program.

2. **File-Based Storage**:
   - All expense data is stored in a JSON file (`expenses.json`). Each time the user adds a new expense, it is saved to this file. When the application starts, it loads any existing data from the file.
   - No in-memory storage is used, ensuring that all data persists even if the program is closed and restarted.

3. **Adding Expenses**:
   - Users can input an expense by providing a description, amount, and date. The expense is added to the file-based storage for future access and reporting.

4. **Viewing Expenses**:
   - The user can view all recorded expenses in a formatted output, listing the date, description, and amount for each entry.

5. **Generating Reports**:
   - Users can generate expense reports for a specified date range. The report shows the total amount spent over that period, helping users track their spending habits.

## Project Structure

```
ExpenseTracker/
├── main.go              # Entry point that runs the menu-driven interface
├── expense.go           # Contains structs and methods for handling expenses
├── storage.go           # Handles reading from and writing to the JSON file
└── report.go            # Logic for generating reports
```

## How to Run the Program

### 1. Clone the Repository

First, clone the repository and navigate to the project directory:

```bash
git clone https://github.com/iovossos/30_Day_Challenge.git
cd 30_Day_Challenge/Day\ 9
```

### 2. Run the Program

Make sure you have Go installed, then run the following command:

```bash
go run main.go
```

### 3. Menu Options

Once the program starts, a menu will appear with the following options:

```
1. Add Expense
2. View Expenses
3. Generate Report
4. Exit
```

- **Add Expense**: You will be prompted to enter a description, amount, and date for the expense. The expense will be saved in the JSON file.
- **View Expenses**: Displays a list of all the expenses stored in the file.
- **Generate Report**: Prompts you for a start and end date, then displays a summary of expenses within that period.
- **Exit**: Closes the application.

### Example Usage

1. **Adding an Expense**:

```
Enter description: Lunch
Enter amount: 12.50
```

The expense will be added and stored in `expenses.json`.

2. **Viewing Expenses**:

Selecting "View Expenses" will display something like:

```
Description: coffee, Amount: 5.00, Date: 2024-10-09
Description: gas, Amount: 25.00, Date: 2024-10-09
```

3. **Generating a Report**:

```
Total expenses: 30
```

### File Breakdown

#### `main.go`
- **`main()`**: This is the entry point that displays the menu and routes user input to the corresponding functionality (e.g., adding expenses, viewing expenses, generating reports).

#### `expense.go`
- **`Expense` struct**: Defines the structure for an expense, including fields for description, amount, and date.
- **Methods** for creating, validating, and handling expense data.

#### `storage.go`
- **`SaveExpenses()`**: Writes all expense data to the JSON file (`expenses.json`).
- **`LoadExpenses()`**: Loads and returns the expense data from the JSON file upon application startup.

#### `report.go`
- **`GenerateReport()`**: Calculates total expenses for a specified date range by filtering the data from the JSON file.

## Features

- **Persistent Storage**: All expenses are saved in a JSON file for future reference, even after the program is closed.
- **Report Generation**: The application can generate a report summarizing expenses.
- **Simple Menu Interface**: Easy-to-use menu system to interact with the expense tracker.

## Limitations
- **No User Management**: The expense data is stored in a single file without any user-specific differentiation.
- **In-Memory Limitations**: Large amounts of data could be slow to process since the entire file is loaded into memory.
- **Basic UI**: The program runs in a command-line interface, and there is no graphical user interface (GUI).

## Future Improvements

1. **Add User Authentication**: Allow multiple users to track their expenses separately.
2. **Improve Reporting**: Add more granular reporting options, such as category-based summaries.
3. **CSV Export**: Add an option to export the expenses to a CSV file for easier analysis.
