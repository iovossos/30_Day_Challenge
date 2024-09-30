
# Currency Converter with External API

This project is a simple currency converter written in Go that fetches real-time exchange rates from an external API. The program takes an amount in one currency and converts it to another currency using live data from ExchangeRate-API or a similar service. The project is built to demonstrate how Go can interact with an external API to fetch and process JSON data.

## How It Works

The program works as follows:

1. **Fetch Exchange Rates**: The program makes an HTTP GET request to fetch the current exchange rates from the external API.
2. **Perform Currency Conversion**: The user provides the amount, source currency, and target currency. The program converts the amount from the source currency to the target currency using the fetched exchange rates.
3. **Display the Result**: The result is printed in the terminal, showing the amount in both the source and target currencies.

## Project Structure

```
Day 4/
├── main.go              # Main file containing all logic for fetching data and performing the conversion
```

## How to Run the Program

### 1. Clone the Repository
Clone this repository and navigate to the `Day 4` directory.

```bash
git clone https://github.com/iovossos/30_Day_Challenge.git
cd 30_Day_Challenge/Day\ 4
```

### 2. Run the Program
Make sure you have Go installed, and then run the following command:

```bash
go run main.go <amount> <from_currency> <to_currency>
```


### Example Usage

Here’s an example of running the program:

```bash
go run main.go 100 USD EUR
```

Expected output:

```
100.00 USD = 84.21 EUR
```

## File Breakdown

### `main.go`

Contains all the logic:

- **`getExchangeRates()`**: Fetches live exchange rates from the external API and parses the JSON response.
- **`main()`**: Takes user input for the amount and currency codes, performs the currency conversion, and prints the result.

