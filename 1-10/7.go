package main

func euler7() {

	//We skip checking the first 6 primes
	numOfPrimes := 6
	lastPrime := 13
	currNum := 14
	for numOfPrimes < 10001 {
		if isPrime(currNum) {
			lastPrime = currNum
			numOfPrimes++
		}
		currNum++
	}

	printAnswer(7, lastPrime)
}
