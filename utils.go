package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

//PrintAnswer -  Prints a formatted answer for problemNum and answer
func PrintAnswer(problemNum int, answer int) {

	fmt.Printf("The answer to Project Euler Problem #%d is :%d\n", problemNum, answer)
}

func PrintBigIntAnswer(problemNum int, answer *big.Int) {

	fmt.Printf("The answer to Project Euler Problem #%d is :%d\n", problemNum, answer)
}

func PrintStringAnswer(problemNum int, answer string) {

	fmt.Printf("The answer to Project Euler Problem #%d is :%s\n", problemNum, answer)
}

func PrintMatrix15(matrix [15][15]int) {

	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Print("\n")
	}
}

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

//IsPalidrome - Checks if num is a palidrome
func IsPalidrome(num int) bool {

	numStr := strconv.Itoa(num)
	revesredNum := Reverse(numStr)

	return strings.Compare(numStr, revesredNum) == 0

}

//Reverse - revesrses string s
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func SumDigits(digits string) int {

	sum := 0
	for i := 0; i < len(digits); i++ {
		digit, _ := strconv.Atoi(string(digits[i]))
		sum += digit
	}

	return sum
}
