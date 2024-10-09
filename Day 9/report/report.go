package report

import (
	"fmt"
	"tracker/expense" // Update to your actual module name
)

// GenerateReport prints a summary report of expenses
func GenerateReport(expenses expense.Expenses) {
	total := 0.0
	for _, e := range expenses {
		total += e.Amount
	}
	fmt.Printf("Total expenses: %.2f\n", total)
}
