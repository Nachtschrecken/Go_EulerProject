package tasks

import "fmt"

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

func MultiplesOf35() {

	cache := make([]int, 0)

	// add all multiples
	for i := 0; i < 1000; i++ {
		if ((i % 3) == 0) || ((i % 5) == 0) {
			cache = append(cache, i)
		}
	}

	// sum up multiples
	sum := 0
	for j := 0; j < len(cache); j++ {
		sum += cache[j]
	}

	fmt.Println(sum)
}
