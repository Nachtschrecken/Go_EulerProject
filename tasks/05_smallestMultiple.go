package tasks

import "fmt"

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func SmallestMultiple() {

	num := smallestMultiple(20)

	fmt.Println(num)
}

func smallestMultiple(k int) int {

	n := k
	for {
		check := 0
		for i := 1; i <= k; i++ {
			if n%i == 0 {
				check++
			}
		}
		if check == k {
			return n
		}
		n++
	}
}
