package main

import (
	"fmt"
	"math"
)

func main() {
	limit := 10000 // Find primes up to this number

	primes := FindPrimes(limit)
	fmt.Println("Primes:", primes)
}

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// FindPrimes finds all prime numbers up to the given limit.
func FindPrimes(limit int) []int {
	var primes []int

	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}
