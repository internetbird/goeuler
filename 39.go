package main

import "github.com/internetbird/goeuler/utils"

func Euler39() {
	max_num_of_solutions := 3
	max_solutions_perimeter := 120

	for a := 1; a < 998; a++ {
		for b := 1; b < 1000-a-1; b++ {

		}
	}

	for p := 121; p <= 1000; p++ {
		num_of_solutions := GetNumOfSolutions(p)
		if num_of_solutions > max_num_of_solutions {
			max_num_of_solutions = num_of_solutions
			max_solutions_perimeter = p
		}
	}

	utils.PrintAnswer(39, max_solutions_perimeter)

}

func GetNumOfSolutions(perimeter int) int {
	return 1
}
