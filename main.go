package main

import (
	"fmt"
	"math"
)

func main() {

	euler1()
	euler2()
	euler6()
	euler10()

}

func euler1() {
	//Project Euler Problem #1
	sum := 0

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	printAnswer(1, sum)

}

func euler2() {
	//Project Euler Problem #2
	prevFib := 1
	currFib := 2
	sum := 0

	for currFib < 4000000 {
		if currFib%2 == 0 {
			sum += currFib
		}

		temp := currFib
		currFib = currFib + prevFib
		prevFib = temp

	}

	printAnswer(2, sum)
}

func euler6() {
	//Project Euler Problem #6
	sumOfSquares := 0
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
		sumOfSquares += int(math.Pow(float64(i), 2))

	}
	squareOfSums := int(math.Pow(float64(sum), 2))

	printAnswer(6, squareOfSums-sumOfSquares)

}

func euler10() {
	//Project Euler Problem #10
	sum := 0

	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	printAnswer(10, sum)

}

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
