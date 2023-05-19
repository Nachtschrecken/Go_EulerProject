package tasks

import "fmt"

func TriangleNumbers() {

	n := 500

	fmt.Println(nMostDivs(n))
}

func nMostDivs(n int) int {
	i := 0
	tNum := 0
	for {
		tNum = 0
		for j := 0; j <= i; j++ {
			tNum += j
		}
		if divisors(tNum) > n {
			return tNum
		}
		fmt.Println(i, tNum, divisors(tNum))
		i++
	}
}

func divisors(n int) int {
	count := 1
	for i := 1; i < n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}
