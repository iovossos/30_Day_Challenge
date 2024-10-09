package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"tracker/expense" // Update to your actual module name
)

const filename = "data/expenses.json"

func main() {
	// Load existing expenses from the file
	expenses, err := expense.LoadExpenses(filename)
	if err != nil {
		log.Fatalf("Error loading expenses: %v", err)
	}

	for {
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. Generate Report")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addExpense(&expenses)
		case 2:
			viewExpenses(expenses)
		case 3:
			generateReport(expenses)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

// Function to add an expense
func addExpense(expenses *expense.Expenses) {
	var desc string
	var amount float64
	fmt.Print("Enter description: ")
	fmt.Scan(&desc)
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)

	newExpense := expense.Expense{
		Description: desc,
		Amount:      amount,
		Date:        time.Now(),
	}

	*expenses = append(*expenses, newExpense)
	if err := expense.SaveExpenses(filename, *expenses); err != nil {
		log.Fatalf("Error saving expenses: %v", err)
	}
	fmt.Println("Expense added successfully.")
}

// Function to view expenses
func viewExpenses(expenses expense.Expenses) {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}

	for _, e := range expenses {
		fmt.Printf("Description: %s, Amount: %.2f, Date: %s\n", e.Description, e.Amount, e.Date.Format("2006-01-02"))
	}
}

// Function to generate a report
func generateReport(expenses expense.Expenses) {
	total := 0.0
	for _, e := range expenses {
		total += e.Amount
	}
	fmt.Printf("Total expenses: %.2f\n", total)
}
