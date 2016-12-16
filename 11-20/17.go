package main

import "strings"

func Euler17() {

	totalLetters := 0

	for i := 1; i <= 1000; i++ {
		numText := ConvertNumToText(i)
		numText = strings.Replace(numText, " ", "", -1)
		totalLetters += len(numText)
	}

	PrintAnswer(17, totalLetters)
}

func ConvertNumToText(num int) string {

	if num == 0 {
		return ""
	} else if num == 1 {
		return "one"
	} else if num == 2 {
		return "two"
	} else if num == 3 {
		return "three"
	} else if num == 4 {
		return "four"
	} else if num == 5 {
		return "five"
	} else if num == 6 {
		return "six"
	} else if num == 7 {
		return "seven"
	} else if num == 8 {
		return "eight"
	} else if num == 9 {
		return "nine"
	} else if num == 10 {
		return "ten"
	} else if num == 11 {
		return "eleven"
	} else if num == 12 {
		return "twelve"
	} else if num == 13 {
		return "thirteen"
	} else if num == 14 {
		return "fourteen"
	} else if num == 15 {
		return "fifteen"
	} else if num == 16 {
		return "sixteen"
	} else if num == 17 {
		return "seventeen"
	} else if num == 18 {
		return "eighteen"
	} else if num == 19 {
		return "nineteen"
	} else if num >= 20 && num < 30 {
		return "twenty " + ConvertNumToText(num-20)
	} else if num >= 30 && num < 40 {
		return "thirty " + ConvertNumToText(num-30)
	} else if num >= 40 && num < 50 {
		return "forty " + ConvertNumToText(num-40)
	} else if num >= 50 && num < 60 {
		return "fifty " + ConvertNumToText(num-50)
	} else if num >= 60 && num < 70 {
		return "sixty " + ConvertNumToText(num-60)
	} else if num >= 70 && num < 80 {
		return "seventy " + ConvertNumToText(num-70)
	} else if num >= 80 && num < 90 {
		return "eighty " + ConvertNumToText(num-80)
	} else if num >= 90 && num < 100 {
		return "ninety " + ConvertNumToText(num-90)
	} else if num == 100 {
		return "one hundred"
	} else if num > 100 && num < 200 {
		return "one hundred and " + ConvertNumToText(num-100)
	} else if num == 200 {
		return "two hundred"
	} else if num > 200 && num < 300 {
		return "two hundred and " + ConvertNumToText(num-200)
	} else if num == 300 {
		return "three hundred"
	} else if num > 300 && num < 400 {
		return "three hundred and " + ConvertNumToText(num-300)
	} else if num == 400 {
		return "four hundred"
	} else if num > 400 && num < 500 {
		return "four hundred and " + ConvertNumToText(num-400)
	} else if num == 500 {
		return "five hundred"
	} else if num > 500 && num < 600 {
		return "five hundred and " + ConvertNumToText(num-500)
	} else if num == 600 {
		return "six hundred"
	} else if num > 600 && num < 700 {
		return "six hundred and " + ConvertNumToText(num-600)
	} else if num == 700 {
		return "seven hundred"
	} else if num > 700 && num < 800 {
		return "seven hundred and " + ConvertNumToText(num-700)
	} else if num == 800 {
		return "eight hundred"
	} else if num > 800 && num < 900 {
		return "eight hundred and " + ConvertNumToText(num-800)
	} else if num == 900 {
		return "nine hundred"
	} else if num > 900 && num < 1000 {
		return "nine hundred and " + ConvertNumToText(num-900)
	} else if num == 1000 {
		return "one thousand"
	}
	return "unknown"
}
