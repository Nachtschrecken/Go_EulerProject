package tasks

import "fmt"

// The sum of the squares of the first ten natural numbers is,
//		1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,
//		(1 + 2 + ... + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is .
//		3025 - 385 = 2640
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

// SOLUTION
// limit = 100
// sum = limit(limit + 1)/2
// sum sq = (2limit + 1)(limit + 1)limit/6
// print sum2 − sum sq

func SumSquareDifference() {

	n := 100

	sum := n * (n + 1) / 2
	sumSq := (2*n + 1) * (n + 1) * n / 6

	result := (sum * sum) - sumSq
	fmt.Println(result)

	/* OLD
	sumOfSquares := 0
	for i := 0; i <= n; i++ {
		sumOfSquares += i * i
	}

	squareOfSums := ((n * (n + 1)) / 2) * ((n * (n + 1)) / 2)

	fmt.Println(squareOfSums - sumOfSquares)
	*/
}
