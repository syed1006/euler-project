package main

import "fmt"

func main() {
	fmt.Println(evenFibonacciSum(4_000_000))
}

func evenFibonacciSum(limit int) int {
	a, b := 1, 2
	sum := 2

	for {
		next := a + b
		if next > limit {
			break
		}

		if next%2 == 0 {
			sum += next
		}

		a, b = b, next
	}

	return sum
}

//4613732
