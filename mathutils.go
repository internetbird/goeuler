package main

import "math"

//IsPrime - Checks is num is a prime number
func IsPrime(num int) bool {

	if num <= 0 {
		return false
	}

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

func GetAbundantNumbers(from int, to int) []int {
	nums := make([]int, 0, to-from)

	for i := from; i <= to; i++ {
		if IsAbundantNumber(i) {
			nums = append(nums, i)
		}
	}
	return nums
}

func IsAbundantNumber(num int) bool {

	properDivisors := ProperDivisorsOfNumber(num)
	divisorsSum := SumIntArray(properDivisors)

	return divisorsSum > num
}

const LIM = 41

var facts [LIM]uint64

func Factorial(n uint64) (res uint64) {
	if facts[n] != 0 {
		res = facts[n]
		return res
	}

	if n > 0 {
		res = n * Factorial(n-1)
		return res
	}

	return 1
}
