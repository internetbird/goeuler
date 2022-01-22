package main

import (
	"fmt"
	"strconv"

	"github.com/internetbird/goeuler/mathutils"
	"github.com/internetbird/goeuler/utils"
)

func Euler37() {
	sumOfPrimes := 0
	numOfFoundPrimes := 0

	for i := 10; i < 10000000; i++ {
		if mathutils.IsPrime(i) {

			primeNumStr := strconv.Itoa(i)
			isTruncablePrime := true

			//check left to right
			for j := 1; j < len(primeNumStr); j++ {

				primeNumSubNumberStr := primeNumStr[j:]
				primeSumNumber, _ := strconv.Atoi(primeNumSubNumberStr)

				if !mathutils.IsPrime(primeSumNumber) {
					isTruncablePrime = false
					break
				}

			}

			if isTruncablePrime {
				//check right to left
				for j := len(primeNumStr) - 1; j > 0; j-- {

					primeNumSubNumberStr := primeNumStr[:j]
					primeSumNumber, _ := strconv.Atoi(primeNumSubNumberStr)

					if !mathutils.IsPrime(primeSumNumber) {
						isTruncablePrime = false
						break
					}

				}

			}

			if isTruncablePrime {
				numOfFoundPrimes++
				fmt.Printf("Found truncable prime  #%d which is :%d\n", numOfFoundPrimes, i)
				sumOfPrimes += i
			}

		}
	}

	utils.PrintAnswer(37, sumOfPrimes)
}
