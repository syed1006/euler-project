/*
<p>If we list all the natural numbers below $10$ that are multiples of $3$ or $5$, we get $3, 5, 6$ and $9$. The sum of these multiples is $23$.</p>
<p>Find the sum of all the multiples of $3$ or $5$ below $1000$.</p>
*/

package main

import "fmt"

func main() {
	limit := 1000
	sum := sumMultiples(3, limit) + sumMultiples(5, limit) - sumMultiples(15, limit)

	fmt.Println(sum)
}

func sumMultiples(divisor, limit int) int {
	k := (limit - 1) / divisor
	return divisor * k * (k + 1) / 2
}

//233168
