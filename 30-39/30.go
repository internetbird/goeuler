package main

import (
	"github.com/internetbird/goeuler/mathutils"
	"github.com/internetbird/goeuler/utils"
)

func Euler30() {

	sum := 0

	for i := 10; i <= 1000000; i++ {
		digits := utils.GetDigits(i)

		powerSum := GetPowerSum(digits)

		if powerSum == i {
			sum += i
		}
	}

	utils.PrintAnswer(30, sum)

}

func GetPowerSum(digits []int) int {

	sum := 0

	for i := 0; i < len(digits); i++ {
		sum += mathutils.Power(digits[i], 5)
	}

	return sum

}
