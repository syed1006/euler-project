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
