package main

import (
	"math"

	"github.com/internetbird/goeuler/mathutils"
	"github.com/internetbird/goeuler/utils"
)

func Euler3() {

	questionNum := 600851475143
	largestPrimeFactor := 3

	topLimit := int(math.Sqrt(float64(questionNum)))

	for i := largestPrimeFactor; i < topLimit; i++ {
		if mathutils.IsPrime(i) && questionNum%i == 0 {
			largestPrimeFactor = i
		}
	}

	utils.PrintAnswer(3, largestPrimeFactor)

}
