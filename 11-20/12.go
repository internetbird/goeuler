package main

import "fmt"

func Euler12() {

	triangle := 1
	counter := 2

	for NumOfDivisors(triangle) <= 500 {
		triangle = triangle + counter
		counter++
	}

	fmt.Printf("\nNum of divisors is :%d\n", NumOfDivisors(triangle))

	PrintAnswer(12, triangle)
}

func NumOfDivisors(num int) int {

	if num == 1 {
		return 1
	}

	if num == 2 {
		return 2
	}

	//The number and 1 are always 2 divisors
	numOfDivisors := 2

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			numOfDivisors++
		}
	}

	return numOfDivisors

}
