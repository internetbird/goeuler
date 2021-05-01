package main

import (
	"github.com/internetbird/goeuler/mathutils"
	"github.com/internetbird/goeuler/utils"
)

func Euler21() {

	sum := 0

	for i := 1; i < 10000; i++ {
		sumOfDivisors := GetSumOfProperDivisors(i)
		sumOfSumOfDivisors := GetSumOfProperDivisors(sumOfDivisors)

		if sumOfSumOfDivisors == i && sumOfDivisors != i {
			sum += i
		}
	}
	utils.PrintAnswer(21, sum)
}

func GetSumOfProperDivisors(num int) int {
	properDivisors := mathutils.ProperDivisorsOfNumber(num)
	sumOfProperDivisors := utils.SumIntArray(properDivisors)

	return sumOfProperDivisors

}
