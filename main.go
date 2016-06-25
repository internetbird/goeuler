package main

import (
	"fmt"
	"math"
)

func main() {

	//	euler1()
	//	euler2()
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

	fmt.Println("The answer to Project Euler Problem #1 is :", sum)
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

	fmt.Println("The answer to Project Euler Problem #2 is :", sum)
}

func euler10() {
	sum := 0

	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println("The answer to Project Euler Problem #10 is :", sum)
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
