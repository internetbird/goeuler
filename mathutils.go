package main

import "math"

//IsPrime - Checks is num is a prime number
func IsPrime(num int) bool {

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

func ProperDivisorsOfNumber(num int) []int {
	divisors := make([]int, 0, num/2)

	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors

}
