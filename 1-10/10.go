package main

func Euler10() {

	sum := 0

	for i := 2; i < 2000000; i++ {
		if IsPrime(i) {
			sum += i
		}
	}

	PrintAnswer(10, sum)
}
