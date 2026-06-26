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
