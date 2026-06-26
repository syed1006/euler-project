/*
<p>The sum of the primes below $10$ is $2 + 3 + 5 + 7 = 17$.</p>
<p>Find the sum of all the primes below two million.</p>
*/

package main

import "fmt"

func main() {
	fmt.Println(sumToPrime(2000000))

	fmt.Println(sumOfInts(primesTillN(2000000)))
}

func sumToPrime(limit int) int {
	sum := 0

	currentPrime := 2

	for {
		sum += currentPrime

		currentPrime = findNextPrimeNumber(currentPrime)

		if currentPrime > limit {
			return sum
		}
	}
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

// 142913828922

// Sieve of Eratosthenes
func primesTillN(limit int) []int {
	prime := make([]bool, limit)

	for i := 2; i < limit; i++ {
		prime[i] = true
	}

	for p := 2; p*p < limit; p++ {

		if prime[p] {
			for multiple := p * p; multiple < limit; multiple += p {
				prime[multiple] = false
			}
		}
	}

	finalPrimes := []int{}

	for index, value := range prime {
		if value {
			finalPrimes = append(finalPrimes, index)
		}
	}

	return finalPrimes
}

func sumOfInts(nums []int) int {

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
