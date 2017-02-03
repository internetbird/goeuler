package main

func Euler21() {

	sum := 0

	for i := 1; i < 10000; i++ {
		sumOfDivisors := GetSumOfProperDivisors(i)
		sumOfSumOfDivisors := GetSumOfProperDivisors(sumOfDivisors)

		if sumOfSumOfDivisors == i && sumOfDivisors != i {
			sum += i
		}
	}
	PrintAnswer(21, sum)
}

func GetSumOfProperDivisors(num int) int {
	properDivisors := ProperDivisorsOfNumber(num)
	sumOfProperDivisors := SumIntArray(properDivisors)

	return sumOfProperDivisors

}
