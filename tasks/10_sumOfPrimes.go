package tasks

import "fmt"

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

func SumOfPrimes() {

	n := 2_000_000
	sum := 0

	for i := 0; i < n; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
