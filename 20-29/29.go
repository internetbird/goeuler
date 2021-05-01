package main

import (
	"math/big"

	"github.com/internetbird/goeuler/mathutils"
	"github.com/internetbird/goeuler/utils"
)

func Euler29() {

	nums := make([]big.Int, 0)

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			power := mathutils.PowerBig(a, b)
			if !utils.ContainsBig(nums, *power) {
				nums = append(nums, *power)
			}
		}
	}
	utils.PrintAnswer(29, len(nums))
}
