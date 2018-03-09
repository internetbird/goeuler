package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func Euler22() {

	sum := 0

	content, err := ioutil.ReadFile("p022_names.txt")

	if err == nil {
		namesTxt := string(content)
		//Clear the quotes
		namesTxt = strings.Replace(namesTxt, "\"", "", -1)
		//Create the names array
		namesArr := strings.Split(namesTxt, ",")
		//Sort the names array
		sort.Strings(namesArr)

		for i := 0; i < len(namesArr); i++ {
			value := GetNameValue(namesArr[i], i+1)
			sum += value
		}

	} else {
		fmt.Println("Error reading names file")
	}

	PrintAnswer(22, sum)
}

func GetNameValue(name string, position int) int {

	stringValue := GetStringValue(name)

	return stringValue * position
}

func GetStringValue(str string) int {

	val := 0

	chars := []rune(str)

	for i := 0; i < len(chars); i++ {
		val += (int(chars[i]) - int('A') + 1)
	}

	return val

}
