package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const apiURL = "https://v6.exchangerate-api.com/v6/YOUR-API-KEY/latest/USD" // Replace with your actual API key

// Define a struct to map the JSON response from the API
type Rates struct {
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

// Fetch the exchange rates from the API
func getExchangeRates() (Rates, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return Rates{}, err
	}
	defer resp.Body.Close()

	// Decode the JSON response into the Rates struct
	var rates Rates
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return Rates{}, err
	}
	return rates, nil
}

func main() {
	// Check if the user provided enough arguments
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <amount> <from_currency> <to_currency>")
		return
	}

	// Parse the arguments (amount to convert, source currency, and target currency)
	amount := os.Args[1]
	fromCurrency := os.Args[2]
	toCurrency := os.Args[3]

	// Fetch the exchange rates from the API
	rates, err := getExchangeRates()
	if err != nil {
		log.Fatalf("Failed to fetch exchange rates: %v", err)
	}

	// Perform the currency conversion
	fromRate, existsFrom := rates.ConversionRates[fromCurrency]
	toRate, existsTo := rates.ConversionRates[toCurrency]

	if !existsFrom || !existsTo {
		fmt.Println("Invalid currency code. Please check the currency symbols.")
		return
	}

	// Parse the amount from string to float64
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		log.Fatalf("Invalid amount: %v", err)
	}

	// Convert from the source currency to USD, then to the target currency
	usdAmount := amountFloat / fromRate
	convertedAmount := usdAmount * toRate

	// Output the conversion result
	fmt.Printf("%.2f %s = %.2f %s\n", amountFloat, fromCurrency, convertedAmount, toCurrency)
}
