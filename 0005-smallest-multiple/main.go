/*
<p>$2520$ is the smallest number that can be divided by each of the numbers from $1$ to $10$ without any remainder.</p>
<p>What is the smallest positive number that is <strong class="tooltip">evenly divisible<span class="tooltiptext">divisible with no remainder</span></strong> by all of the numbers from $1$ to $20$?</p>
*/

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
