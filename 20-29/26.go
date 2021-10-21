package main

import (
	"fmt"
	"math/big"
	"strings"
)

func Euler26() {

	for i := 2; i < 1000; i++ {

		nominator := big.NewFloat(1)
		denominator := big.NewFloat(float64(i))

		fraction := new(big.Float).Quo(nominator, denominator)
		fractionString := fraction.Text('f', 100)

		fractionPart := fractionString[2:]

		for j := 9; j < len(fractionPart)/2; j++ {

			zerosString := strings.Repeat("0", j)

			for offset := 0; offset < len(fractionPart)-j*2; offset++ {
				part1 := fractionPart[offset : offset+j]
				part2 := fractionPart[offset+j : offset+j*2]

				if part1 == part2 && part1 != zerosString {
					fmt.Printf("\n -- 1/%d = %s", i, fractionString)
					fmt.Printf("\nFraction part is: %s", fractionPart)
					fmt.Printf("\nPart 1: %s Part 2: %s", part1, part2)
				}

			}
		}

	}

}
