package main

import (
	"math/big"

	"github.com/internetbird/goeuler/utils"
)

func Euler16() {

	powVal := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)

	sum := utils.SumDigits(powVal.String())

	utils.PrintAnswer(16, sum)
}
