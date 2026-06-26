/*
<p>By listing the first six prime numbers: $2, 3, 5, 7, 11$, and $13$, we can see that the $6$th prime is $13$.</p>
<p>What is the $10\,001$st prime number?</p>
*/

package main

import "fmt"

func main() {
	fmt.Println(primeAtPosition(10001))
}

func primeAtPosition(targetPosition int) int {

	currentPosition := 1
	currentPrime := 2

	for currentPosition < targetPosition {
		currentPrime = findNextPrimeNumber(currentPrime)
		currentPosition++
	}

	return currentPrime

}

func findNextPrimeNumber(number int) int {
	if number == 2 {
		return 3
	}

	for {
		number += 2
		if isPrime(number) {
			return number
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// 104743
