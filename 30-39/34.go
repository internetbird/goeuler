package main

import (
	"github.com/internetbird/goeuler/mathutils"
	"github.com/internetbird/goeuler/utils"
)

func Euler34() {

	sum := 0

	for i := 10; i < 10000000; i++ {
		if IsDigitalFactorial(i) {
			sum += i
		}
	}

	utils.PrintAnswer(34, sum)
}

func IsDigitalFactorial(num int) bool {
	digits := utils.GetDigits(num)

	factorialSum := 0

	for i := 0; i < len(digits); i++ {
		factorialSum += int(mathutils.Factorial(uint64(digits[i])))
	}

	return factorialSum == num
}
