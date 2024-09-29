
# Prime Number Finder

This project is a simple program written in Go that finds all prime numbers up to a specified limit. It contains a single function that checks for prime numbers and returns them as a list. The solution does not use concurrency or workers, making it simple and easy to understand.

## How It Works

The program consists of two main functions:

1. **`isPrime`**: This function checks if a given number is prime.
2. **`FindPrimes`**: This function finds all prime numbers between 2 and a given `limit`, using the `isPrime` function.

The **`isPrime`** function works by testing divisibility of a number by all integers between 2 and the square root of the number. If the number is divisible by any of these, it is not a prime.

The **`FindPrimes`** function loops through all numbers from 2 to the given `limit`. For each number, it checks if it's prime using `isPrime`. If it’s prime, it appends the number to a list of primes.

## Project Structure

```
prime-finder/
├── main.go              # Main file that contains all functions (isPrime, FindPrimes) and runs the program
```

## How to Run the Program

### 1. Clone the Repository
Clone this repository and navigate to the `Day 2` directory.

```bash
git clone https://github.com/iovossos/30_Day_Challenge.git
cd 30_Day_Challenge/Day\ 2
```
### 2. Run the Program
Make sure you have Go installed, and then run the following command:

```bash
go run main.go
```

### Example Usage

You can specify the limit up to which you want to find prime numbers. For example, to find all primes up to 100, the program can be run like this:

```go
package main

import (
    "fmt"
)

func main() {
    limit := 100  // Set the limit for prime number search
    primes := FindPrimes(limit)
    fmt.Println("Primes:", primes)
}

```

### Output

If you run the program with a limit of 100, the output will be:

```
Primes: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
```

## File Breakdown

### `main.go`

Contains all the logic:

- **`isPrime(n int) bool`**: Checks if a given number `n` is a prime number.
- **`FindPrimes(limit int) []int`**: Finds all prime numbers up to the given `limit`.
