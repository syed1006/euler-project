package main

import "fmt"

func main() {
	fmt.Println(smallestMultiple(20))
}

func smallestMultiple(limit int) int {
	number := limit + 1
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
