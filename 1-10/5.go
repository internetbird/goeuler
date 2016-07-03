package main

func euler5() {
	smallestMultiple := 2520

	for !devides1to20(smallestMultiple) {
		smallestMultiple += 20
	}

	printAnswer(5, smallestMultiple)

}

func devides1to20(num int) bool {
	for i := 3; i < 20; i++ {
		if num%i != 0 {
			return false
		}
	}

	return true
}
