package tasks

import (
	"fmt"
	"math"
)

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func LargestPrimeFactor() {

	n := 600851475143

	largest := checkPrime(n)

	fmt.Println(largest)
}

func checkPrime(n int) int {
	i := 1
	for {
		if isPrime(i) {
			if isPrime(n) {
				return n
			}
			if n%i == 0 {
				n = n / i
			}
		}
		i++
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
