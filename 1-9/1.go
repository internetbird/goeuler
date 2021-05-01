package main

import "github.com/internetbird/goeuler/utils"

func Euler1() {
	//Project Euler Problem #1
	sum := 0

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	utils.PrintAnswer(1, sum)

}
