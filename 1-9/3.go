package main

import (
	"math"
)

func Euler3() {

	questionNum := 600851475143
	largestPrimeFactor := 3

	topLimit := int(math.Sqrt(float64(questionNum)))

	for i := largestPrimeFactor; i < topLimit; i++ {
		if IsPrime(i) && questionNum%i == 0 {
			largestPrimeFactor = i
		}
	}

	PrintAnswer(3, largestPrimeFactor)

}
