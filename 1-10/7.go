package main

func Euler7() {

	//We skip checking the first 6 primes
	numOfPrimes := 6
	lastPrime := 13
	currNum := 14
	for numOfPrimes < 10001 {
		if IsPrime(currNum) {
			lastPrime = currNum
			numOfPrimes++
		}
		currNum++
	}

	PrintAnswer(7, lastPrime)
}
