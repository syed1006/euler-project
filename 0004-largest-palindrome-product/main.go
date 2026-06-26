/*
<p>A palindromic number reads the same both ways. The largest palindrome made from the product of two $2$-digit numbers is $9009 = 91 \times 99$.</p>
<p>Find the largest palindrome made from the product of two $3$-digit numbers.</p>
*/

package main

import "fmt"

func main() {
	fmt.Println(findPalindrome())
}

func findPalindrome() int {

	palindrome := 0

	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := i * j

			if product <= palindrome {
				break
			}

			if isPalindrome(product) {
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

	for number > 0 {

		reverse = reverse*10 + number%10

		number = number / 10

	}

	return reverse
}

// 906609
