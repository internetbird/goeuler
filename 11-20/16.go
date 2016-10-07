package main

import (
	"math/big"
	"strconv"
)

func Euler16() {

	powVal := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)

	powValStr := powVal.String()

	sum := 0

	for i := 0; i < len(powValStr); i++ {
		digit, _ := strconv.Atoi(string(powValStr[i]))
		sum += digit
	}

	PrintAnswer(16, sum)
}
