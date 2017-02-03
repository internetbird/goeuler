package main

import (
	"math/big"
	"strconv"
)

func Euler16() {

	powVal := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)

	powValStr := powVal.String()

	sum := SumDigits(powVal.String())

	PrintAnswer(16, sum)
}
