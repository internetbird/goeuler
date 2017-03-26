package main

import "math/big"

func Euler29() {

	nums := make([]big.Int, 0)

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			power := PowerBig(a, b)
			if !ContainsBig(nums, *power) {
				nums = append(nums, *power)
			}
		}
	}
	PrintAnswer(29, len(nums))
}
