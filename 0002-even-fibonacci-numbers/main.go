package main

import "fmt"

func main() {

	fibonacciSeries := []int{1, 2}

	evenSum := generateFibonacci(1, 2, 4000000, &fibonacciSeries, 2)

	fmt.Println(evenSum)
}

func generateFibonacci(firstTerm int, secondTerm int, limit int, series *[]int, sum int) int {
	nextTerm := firstTerm + secondTerm

	if nextTerm%2 == 0 {
		sum += nextTerm
	}

	if nextTerm > limit {
		return sum
	}

	*series = append(*series, nextTerm)

	return generateFibonacci(secondTerm, nextTerm, limit, series, sum)
}

//4613732
