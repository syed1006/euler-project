package main

import "fmt"

func main() {
	n := 100

	sumOfSquares := sumOfSquareOfNaturalNumbers(n)

	sumOfNumbers := sumOfNaturalNumbers(100)

	result := sumOfNumbers*sumOfNumbers - sumOfSquares

	fmt.Println(result)

}

func sumOfNaturalNumbers(n int) int {

	return (n * (n + 1)) / 2
}

func sumOfSquareOfNaturalNumbers(n int) int {

	return (n * (n + 1) * (2*n + 1)) / 6
}

// 25164150
