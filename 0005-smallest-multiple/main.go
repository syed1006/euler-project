package main

import "fmt"

func main() {
	fmt.Println(smallestMultipleUsingLcm(20))
}

func smallestMultiple(limit int) int {
	number := limit
	for {
		if evenlyDivisible(number, limit) {
			return number
		}
		number += limit
	}
}

func evenlyDivisible(number int, limit int) bool {
	for i := limit; i >= 2; i-- {
		if number%i != 0 {
			return false
		}
	}

	return true
}

// 232792560

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func leastCommonMultiple(a, b int) int {
	return a / greatestCommonDivisor(a, b) * b
}

func smallestMultipleUsingLcm(limit int) int {
	result := 1

	for i := 2; i <= limit; i++ {
		result = leastCommonMultiple(result, i)
	}

	return result
}

// LCM(1,2,3,...,20)
