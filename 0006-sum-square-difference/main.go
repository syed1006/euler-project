/*
<p>The sum of the squares of the first ten natural numbers is,</p>
$$1^2 + 2^2 + ... + 10^2 = 385.$$
<p>The square of the sum of the first ten natural numbers is,</p>
$$(1 + 2 + ... + 10)^2 = 55^2 = 3025.$$
<p>Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is $3025 - 385 = 2640$.</p>
<p>Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.</p>
*/

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
