package main

import (
	"fmt"

	"github.com/internetbird/goeuler/utils"
)

//Euler 33 - Solves Project Euler problem #33
func Euler33() {
	for i := 10; i < 98; i++ { //numerator
		for j := i + 1; j < 100; j++ { //denumerator
			result := float64(i) / float64(j)

			canceledNumerator := 0
			canceledDenumerator := 0

			numeratorDigits := utils.GetDigits(i)
			denumeratorDigits := utils.GetDigits(j)

			searchIndex0 := utils.FindIndex(denumeratorDigits, numeratorDigits[0])
			searchIndex1 := utils.FindIndex(denumeratorDigits, numeratorDigits[1])

			if searchIndex0 != -1 {
				canceledNumerator = numeratorDigits[1]
				canceledDenumerator = denumeratorDigits[GetOtherIndex(searchIndex0)]

			} else if searchIndex1 != -1 && numeratorDigits[1] != 0 { //Check that it's not trivial (for example 30/50 = 3/5)
				canceledNumerator = numeratorDigits[0]
				canceledDenumerator = denumeratorDigits[GetOtherIndex(searchIndex1)]
			}

			if canceledNumerator != 0 { //If there is one common digits

				cancledResult := float64(canceledNumerator) / float64(canceledDenumerator)

				if result == cancledResult {
					fmt.Printf("%d/%d = %f\n", i, j, result)
					fmt.Printf("%d/%d = %f\n", canceledNumerator, canceledDenumerator, cancledResult)
				}
			}

		}
	}
}

//GetOtherIndex - returns the remaining int index after cancelling
func GetOtherIndex(index int) int {
	if index == 0 {
		return 1
	}

	return 0
}
