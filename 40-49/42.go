package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Euler42() {

	numOfTriangleWords := 0

	content, err := ioutil.ReadFile("p042_words.txt")

	if err == nil {
		wordsTxt := string(content)
		//Clear the quotes
		wordsTxt = strings.Replace(wordsTxt, "\"", "", -1)
		//Create the names array
		wordsArr := strings.Split(wordsTxt, ",")

		for i := 0; i < len(wordsArr); i++ {
			wordValue := GetWordValue(wordsArr[i])

			if IsTriangleNumber(wordValue) {
				fmt.Printf("Word %s with value %d is triangle\n", wordsArr[i], wordValue)
				numOfTriangleWords++
			}

		}

	} else {
		fmt.Println("Error reading names file")
	}

	PrintAnswer(42, numOfTriangleWords)
}

func IsTriangleNumber(num int) bool {
	for i := 1; i < 10000; i++ {

		curr_triangle_num := i * (i + 1) / 2

		if curr_triangle_num == num {
			return true
		}

		if curr_triangle_num > num {
			return false
		}
	}

	return false

}

func GetWordValue(word string) int {
	val := 0

	chars := []rune(word)

	for i := 0; i < len(chars); i++ {
		val += (int(chars[i]) - int('A') + 1)
	}

	return val

}
