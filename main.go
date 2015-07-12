package main

import (
	"fmt"
)

func main() {

	euler1()
	euler2()
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

	fmt.Println("The answer to Project Euler Problem #22 is :", sum)
}
