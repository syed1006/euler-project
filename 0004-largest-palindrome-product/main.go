package main

import "fmt"

func main() {
	fmt.Println(findPalindrome())
}

func findPalindrome() int {

	palindrome := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := i * j

			if isPalindrome(product) && product > palindrome {
				palindrome = product
			}
		}
	}

	return palindrome
}

func isPalindrome(number int) bool {
	return number == reverseNumber(number)
}

func reverseNumber(number int) int {
	reverse := 0

	for {

		reverse = reverse*10 + number%10

		number = number / 10

		if number <= 0 {
			break
		}
	}

	return reverse
}

// 906609
