package main

func euler10() {
	//Project Euler Problem #10
	sum := 0

	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	printAnswer(10, sum)

}
