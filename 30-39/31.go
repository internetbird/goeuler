package main

import (
	"fmt"
)

func Euler31() {

	//Get the num of ways to pay 2 pounds

	for i := 1; i < 10; i++ {
		numOfWays := GetNumOfWays(i)
		fmt.Printf("Num of ways for #%d is :%d\n", i, numOfWays)
	}

}

func GetNumOfWays(amount int) int {

	numOfWays := 0

	fullCoinBonus := 0

	//for 1p there is only one option
	if amount == 1 {
		return 1
	}

	if amount%2 == 0 {
		fullCoinBonus++
	}

	if amount%5 == 0 {
		fullCoinBonus++
	}

	if amount > 5 {
		fullCoinBonus++
	}

	if amount%10 == 0 {
		fullCoinBonus++
	}

	if amount > 10 {
		fullCoinBonus++
	}

	if amount%20 == 0 {
		fullCoinBonus++
	}

	if amount > 20 {
		fullCoinBonus++
	}

	if amount%50 == 0 {
		fullCoinBonus++
	}

	if amount > 50 {
		fullCoinBonus++
	}

	if amount%100 == 0 {
		fullCoinBonus++
	}

	if amount%200 == 0 {
		fullCoinBonus++
	}

	//fmt.Printf("fullCoinBonus is #%d\n", fullCoinBonus)

	numOfWays = GetNumOfWays(amount-1) + fullCoinBonus

	return numOfWays
}
