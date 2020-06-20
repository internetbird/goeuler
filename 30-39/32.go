package main

func Euler32() {

	special_products := make([]int, 0)

	//Product of xx * yyy = zzzz
	for i := 12; i <= 98 ; i++ {
		for j := 123 ; j <= 987; j++ {

			product := i*j

			if product > 9999 {
				continue
			}

			mul1Digits := GetDigits(i)
			mul2Digits := GetDigits(j)
			productDigits := GetDigits(product)

			if CheckAllDigits(mul1Digits, mul2Digits, productDigits) {
				if  !Contains(special_products, product) {
					special_products = append(special_products, product)
				}
			}
		}
	}

	//Product of x * yyyy = zzzz
	for i := 1; i < 10 ; i++ {
		for j := 1000; j < 10000; j++ {
			product := i*j

			if product > 9999 {
				continue
			}

			mul1Digits := GetDigits(i)
			mul2Digits := GetDigits(j)
			productDigits := GetDigits(product)

			if CheckAllDigits(mul1Digits, mul2Digits, productDigits) {
				if  !Contains(special_products, product) {
					special_products = append(special_products, product)
				}
			}

		}
	}
	PrintAnswer(32, SumIntArray(special_products))
}

func CheckAllDigits(mult1Digits []int, mult2Digits []int, productDigits []int) bool {
	
	for i := 1; i < 10; i++ {
		if !Contains(mult1Digits, i) && !Contains(mult2Digits, i) && !Contains(productDigits, i) {
			return false
		}
	}

	return true
}