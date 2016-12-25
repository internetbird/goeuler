package main

import "math/big"

func Euler20() {

	//sumOfDigits := 0
	factorialWinthoutZeros := big.NewInt(1)

	for i := 2; i <= 99; i++ {

		multiplier := big.NewInt(int64(i))
		factorialWinthoutZeros.Mul(factorialWinthoutZeros, multiplier)
	}

	sum := SumDigits(factorialWinthoutZeros.String())

	PrintAnswer(20, sum)
}
