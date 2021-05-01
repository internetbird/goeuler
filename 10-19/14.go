package main

import "github.com/internetbird/goeuler/utils"

func Euler14() {

	maxChainLengh := 0
	maxChainNum := 21

	for i := 21; i < 1000000; i++ {
		currLengh := GetCollatzSequenceLength(i)

		if currLengh > maxChainLengh {
			maxChainLengh = currLengh
			maxChainNum = i
		}
	}

	utils.PrintAnswer(14, maxChainNum)
}

func GetCollatzSequenceLength(num int) int {
	length := 0

	for num != 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num*3 + 1
		}
		length++
	}
	return length
}
