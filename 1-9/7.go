package main

import (
	"github.com/internetbird/goeuler/mathutils"	"github.com/internetbird/goeuler/utils"
)


func Euler7() {

	//We skip checking the first 6 primes
	numOfPrimes := 6
	lastPrime := 13
	currNum := 14
	for numOfPrimes < 10001 {
		if mathutils.IsPrime(currNum) {
			lastPrime = currNum
			numOfPrimes++
		}
		currNum++
	}

	utils.PrintAnswer(7, lastPrime)
}
