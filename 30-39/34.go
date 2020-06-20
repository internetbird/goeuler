package main

func Euler34() {

	sum := 0

	for i := 10; i < 10000000; i++ {
		if IsDigitalFactorial(i) {
			sum += i
		}
	}

	PrintAnswer(34, sum)
}

func IsDigitalFactorial(num int) bool {
	digits := GetDigits(num)

	factorialSum := 0

	for i := 0; i < len(digits); i++ {
		factorialSum += int(Factorial(uint64(digits[i])))
	}

	return factorialSum == num
}
