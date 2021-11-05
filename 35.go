package main

import (
	"github.com/internetbird/goeuler/mathutils"
	"github.com/internetbird/goeuler/utils"
)

func Euler35() {
	numOfCiruculerPrimes := 0

	for i := 1; i < 1000000; i++ {
		if mathutils.IsPrime(i) {

		}
	}

	utils.PrintAnswer(35, numOfCiruculerPrimes)
}
