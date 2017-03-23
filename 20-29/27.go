package main

import "math"
import "fmt"

func Euler27() {

	max_primes := 0
	max_a := -99
	max_b := -1000

	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {

			if !IsPrime(b) {
				continue
			}

			n := 1
			//value for n=1
			value := 1 + a + b

			for IsPrime(value) {

				if n > max_primes {
					max_primes = n
					max_a = a
					max_b = b
				}

				n = n + 1
				value = int(math.Pow(float64(n), float64(2))) + a*n + b
			}
		}
	}

	fmt.Printf("max a is %d\nmax b is %d\nn = %d\nproduct=%d ", max_a, max_b, max_primes, max_a*max_b)

}
