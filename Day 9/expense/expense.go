package expense

import (
	"encoding/json"
	"os"
	"time"
)

// Expense represents an expense with a description, amount, and date
type Expense struct {
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

// Expenses is a slice of Expense
type Expenses []Expense

// LoadExpenses loads expenses from a JSON file
func LoadExpenses(filename string) (Expenses, error) {
	var expenses Expenses

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return expenses, nil // Return empty expenses if file does not exist
	}

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return expenses, err
	}
	defer file.Close()

	// Decode the JSON data
	if err := json.NewDecoder(file).Decode(&expenses); err != nil {
		return expenses, err
	}
	return expenses, nil
}

// SaveExpenses saves expenses to a JSON file
func SaveExpenses(filename string, expenses Expenses) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(expenses)
}
