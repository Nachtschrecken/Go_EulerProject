package tasks

import (
	"fmt"
	"math/big"
)

func LatticePaths() {

	k := int64(20)

	// Calculate the binomial coefficient
	numerator := factorial(2 * k)
	denominator := new(big.Int).Mul(factorial(k), factorial(2*k-k))
	coeff := new(big.Int).Div(numerator, denominator)

	fmt.Println(coeff)
}

func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}