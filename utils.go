package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func PrintAnswer(problemNum int, answer int) {
	fmt.Printf("The answer to Project Euler Problem #%d is :%d\n", problemNum, answer)
}

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

func IsPalidrome(num int) bool {

	numStr := strconv.Itoa(num)
	revesredNum := Reverse(numStr)

	return strings.Compare(numStr, revesredNum) == 0

}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
