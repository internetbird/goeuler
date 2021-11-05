package main

import (
	"strconv"

	"github.com/internetbird/goeuler/utils"
)

func Euler36() {
	sumOfNumbers := 0

	for i := 0; i < 1000000; i++ {
		if utils.IsPalidrome(i) {
			//convert to binary
			base2Str := strconv.FormatInt(int64(i), 2)

			if utils.IsPalidromeString(base2Str) {
				sumOfNumbers += i

			}
		}
	}

	utils.PrintAnswer(36, sumOfNumbers)
}
