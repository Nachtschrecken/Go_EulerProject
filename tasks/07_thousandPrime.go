package tasks

import "fmt"

func ThousandstPrime() {

	n := 10_001

	fmt.Println(calculate(n))
}

func calculate(n int) int {
	count := 0
	i := 1
	for {
		if isPrime(i) {
			count++
		}
		if count == n {
			return i
		}
		i++
	}
}
