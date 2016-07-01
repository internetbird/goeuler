package main

import (
	"fmt"
	"math"
)

func printAnswer(problemNum int, answer int) {
	fmt.Printf("The answer to Project Euler Problem #%d is :%d\n", problemNum, answer)
}

func isPrime(num int) bool {

	if num == 2 || num == 3 {
		return true
	}
	checkLimit := int(math.Sqrt(float64(num)))

	for i := 2; i <= checkLimit; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
