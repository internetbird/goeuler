package main

import (
	"github.com/internetbird/goeuler/mathutils"
	"github.com/internetbird/goeuler/utils"
)

func Euler10() {

	sum := 0

	for i := 2; i < 2000000; i++ {
		if mathutils.IsPrime(i) {
			sum += i
		}
	}

	utils.PrintAnswer(10, sum)
}
