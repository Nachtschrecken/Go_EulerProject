package tasks

import (
	"fmt"
	"math/big"
	"strconv"
)

func PowerDigitSum() {

	base := big.NewInt(2)
	expo := big.NewInt(1000)
	power := new(big.Int)
	power.Exp(base, expo, nil)

	str := power.String()
	sum := 0

	for _, ch := range str {
		digit, _ := strconv.Atoi(string(ch))
		sum += digit
	}

	fmt.Println(sum)
}
