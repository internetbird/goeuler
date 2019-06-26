package main

import (
	"fmt"
)

func Euler26() {
	for i := 1; i < 20; i++ {
		fmt.Printf("\n -- 1/%d = %.20f", i, 1/float64(i))
	}

}
