package main

func Euler5() {
	smallestMultiple := 2520

	for !Devides1to20(smallestMultiple) {
		smallestMultiple += 20
	}

	PrintAnswer(5, smallestMultiple)

}

func Devides1to20(num int) bool {
	for i := 3; i < 20; i++ {
		if num%i != 0 {
			return false
		}
	}

	return true
}
