package main

import (
	"strconv"

	"github.com/internetbird/goeuler/utils"
)

func Euler38() {

	maxPanDigital := 0

	for num := 2; num < 100000; num++ {

		concatinatedProduct := strconv.Itoa(num)
		multiplier := 2

		for len(concatinatedProduct) < 9 {
			product := num * multiplier
			concatinatedProduct += strconv.Itoa(product)
			multiplier++
		}

		if utils.IsPanDigital(concatinatedProduct) {

			productValue, _ := strconv.Atoi(concatinatedProduct)
			if productValue > maxPanDigital {
				maxPanDigital = productValue
			}
		}
	}

	utils.PrintAnswer(38, maxPanDigital)

}
