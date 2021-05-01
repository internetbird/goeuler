package main

import "github.com/internetbird/goeuler/utils"

func Euler32() {

	special_products := make([]int, 0)

	//Product of xx * yyy = zzzz
	for i := 12; i <= 98; i++ {
		for j := 123; j <= 987; j++ {

			product := i * j

			if product > 9999 {
				continue
			}

			mul1Digits := utils.GetDigits(i)
			mul2Digits := utils.GetDigits(j)
			productDigits := utils.GetDigits(product)

			if CheckAllDigits(mul1Digits, mul2Digits, productDigits) {
				if !utils.Contains(special_products, product) {
					special_products = append(special_products, product)
				}
			}
		}
	}

	//Product of x * yyyy = zzzz
	for i := 1; i < 10; i++ {
		for j := 1000; j < 10000; j++ {
			product := i * j

			if product > 9999 {
				continue
			}

			mul1Digits := utils.GetDigits(i)
			mul2Digits := utils.GetDigits(j)
			productDigits := utils.GetDigits(product)

			if CheckAllDigits(mul1Digits, mul2Digits, productDigits) {
				if !utils.Contains(special_products, product) {
					special_products = append(special_products, product)
				}
			}

		}
	}
	utils.PrintAnswer(32, utils.SumIntArray(special_products))
}

func CheckAllDigits(mult1Digits []int, mult2Digits []int, productDigits []int) bool {

	for i := 1; i < 10; i++ {
		if !utils.Contains(mult1Digits, i) && !utils.Contains(mult2Digits, i) && !utils.Contains(productDigits, i) {
			return false
		}
	}

	return true
}
