package main

import (
	"fmt"
	"strconv"

	"github.com/internetbird/goeuler/utils"
)

func Euler43() {

	sumOfNums := 0

	for i := 1000000000; i < 9999999999; i++ {
		if utils.IsFullPanDigital(i) {

			numStr := strconv.Itoa(i)

			subNum1, _ := strconv.Atoi(numStr[1:3])
			subNum2, _ := strconv.Atoi(numStr[2:4])
			subNum3, _ := strconv.Atoi(numStr[3:5])
			subNum4, _ := strconv.Atoi(numStr[4:6])
			subNum5, _ := strconv.Atoi(numStr[5:7])
			subNum6, _ := strconv.Atoi(numStr[6:8])
			subNum7, _ := strconv.Atoi(numStr[7:])

			if (subNum1%2 == 0) && (subNum2%3 == 0) && (subNum3%5 == 0) && (subNum4%7 == 0) && (subNum5%11 == 0) && (subNum6%13 == 0) && (subNum7%17 == 0) {

				fmt.Printf("The pandigital num: %d has the property", i)
				sumOfNums += i

			}

		}
	}

	utils.PrintAnswer(43, sumOfNums)
}
