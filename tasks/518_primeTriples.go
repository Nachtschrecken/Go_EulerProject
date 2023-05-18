package tasks

import (
	"fmt"
)

/*
Let S(n) = Σ a+b+c over all triples (a,b,c) such that:

a, b, and c are prime numbers.
a < b < c < n.
a+1, b+1, and c+1 form a geometric sequence.
For example, S(100) = 1035 with the following triples:

(2, 5, 11), (2, 11, 47), (5, 11, 23), (5, 17, 53), (7, 11, 17), (7, 23, 71), (11, 23, 47), (17, 23, 31), (17, 41, 97), (31, 47, 71), (71, 83, 97)

Find S(108)
*/

func PrimeTriplesGeo() {

	LIMIT := int(100_000_000)

	var ans int
	isPrimeList := make([]bool, LIMIT)
	for i := range isPrimeList {
		isPrimeList[i] = isPrime(i)
	}

	for x := 1; x <= LIMIT/4; x++ {
		for y := 1; ; y++ {
			a := x*y*y - 1
			if a >= LIMIT {
				break
			}
			if !isPrimeList[a] {
				continue
			}
			for z := y + 1; ; z++ {
				if gcd(y, z) != 1 {
					continue
				}
				c := x*z*z - 1
				if c >= LIMIT {
					break
				}
				if isPrimeList[c] {
					b := x*y*z - 1
					if isPrimeList[b] {
						ans += a + b + c
					}
				}
			}
		}
	}

	fmt.Println(ans)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
