package main

import "fmt"

func main() {
	fmt.Println(findLargestPrimeFactor(600851475143, 2))
	fmt.Println(largestPrimeFactor(600851475143))
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

	if number < 2 {
		return false
	}

	for i := 2; i*i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func largestPrimeFactor(n int64) int64 {
	var factor int64 = 2

	for n%2 == 0 {
		n /= 2
	}

	factor = 3

	for factor*factor <= n {
		if n%factor == 0 {
			n /= factor
		} else {
			factor += 2
		}
	}

	if n > 1 {
		return n
	}

	return factor
}

// 6857
