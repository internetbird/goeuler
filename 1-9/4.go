package main

import (
	"github.com/internetbird/goeuler/utils"
)

func Euler4() {

	largest_palidrome := 0

	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {

			mult := i * j

			if utils.IsPalidrome(mult) && mult > largest_palidrome {
				largest_palidrome = mult
			}
		}
	}

	utils.PrintAnswer(4, largest_palidrome)
}
