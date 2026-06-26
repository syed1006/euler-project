package main

import "fmt"

func main() {
	fmt.Println(findLargestPrimeFactor(600851475143, 2))
}

func findLargestPrimeFactor(number int, primeDivisor int) int {

	for number%primeDivisor == 0 {
		number = number / primeDivisor
	}

	if number == 1 {
		return primeDivisor
	}

	return findLargestPrimeFactor(number, findNextPrimeNumber(primeDivisor))
}

func findNextPrimeNumber(number int) int {
	nextNumber := number + 1

	if isPrime(nextNumber) {
		return nextNumber
	}

	return findNextPrimeNumber(nextNumber)

}

func isPrime(number int) bool {

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

// 6857
