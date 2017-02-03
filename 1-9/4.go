package main

func Euler4() {

	largest_palidrome := 0

	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {

			mult := i * j

			if IsPalidrome(mult) && mult > largest_palidrome {
				largest_palidrome = mult
			}
		}
	}

	PrintAnswer(4, largest_palidrome)
}
