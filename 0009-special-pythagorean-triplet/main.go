/*
<p>A Pythagorean triplet is a set of three natural numbers, $a \lt b \lt c$, for which,
$$a^2 + b^2 = c^2.$$</p>
<p>For example, $3^2 + 4^2 = 9 + 16 = 25 = 5^2$.</p>
<p>There exists exactly one Pythagorean triplet for which $a + b + c = 1000$.<br>Find the product $abc$.</p>

*/

package main

import "fmt"

func main() {

	stop := false

	for m := 2; ; m++ {
		for n := 1; n < m; n++ {
			a, b, c := generatePythagoReanTriplets(m, n)

			if a+b+c == 1000 {
				stop = true
				fmt.Println(a, b, c)
				fmt.Println(a * b * c)
				break
			}
		}

		if stop {
			break
		}
	}
}

func generatePythagoReanTriplets(m int, n int) (int, int, int) {
	return m*m - n*n, 2 * m * n, m*m + n*n
}

// 375 200 425
// 31875000
