package main

import (
	"github.com/internetbird/goeuler/mathutils"
	"github.com/internetbird/goeuler/utils"
)

func Euler23() {
	sum := 0

	abundantNumbers := mathutils.GetAbundantNumbers(12, 28123)

	for i := 1; i < 28124; i++ {
		if i < 24 || !IsAbundantSum(i, abundantNumbers) {
			sum += i
		}
	}

	utils.PrintAnswer(23, sum)
}

func IsAbundantSum(num int, abundantNumbers []int) bool {
	isAbundantSum := false

	for i := 0; i < len(abundantNumbers); i++ {

		currAbundantNum := abundantNumbers[i]

		if currAbundantNum >= num {
			break
		}

		complementSum := num - currAbundantNum

		if complementSum == currAbundantNum || utils.Contains(abundantNumbers, complementSum) {
			isAbundantSum = true
			break
		}

	}
	return isAbundantSum
}
