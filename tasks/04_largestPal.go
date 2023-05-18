package tasks

import (
	"fmt"
	"sort"
	"strconv"
)

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func LargestPalindromeProduct() {

	pals := make([]int, 0)

	for i := 1000; i > 0; i-- {
		for j := 1000; j > 0; j-- {

			str := strconv.Itoa(i * j)
			runes := []rune(str)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			reversedStr := string(runes)

			if str == reversedStr {
				pals = append(pals, i*j)
			}
		}
	}

	sort.Ints(pals)
	fmt.Println(pals[len(pals)-1])
}
