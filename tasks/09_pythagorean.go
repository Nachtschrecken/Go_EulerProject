package tasks

import (
	"fmt"
	"math"
)

/* A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
	a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 5^2 = 9 + 16 = 25.
There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc. */

func SpecialPythagoreanTriplet() {

	n := 1000
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			c := math.Sqrt(float64(a*a + b*b))
			if (c == float64(int(c))) && a < b && b < int(c) {
				if a+b+int(c) == n {
					fmt.Println(a * b * int(c))
				}
			}
		}
	}
}
