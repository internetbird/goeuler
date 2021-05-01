package main

import "github.com/internetbird/goeuler/utils"

func Euler19() {

	numOfFirstMonthSundays := 0

	for year := 1901; year < 2001; year++ {
		for month := 1; month <= 12; month++ {
			firstMonthDay := GetFirstMonthDay(month, year)
			if firstMonthDay == 0 { //Sunday
				numOfFirstMonthSundays++
			}
		}
	}
	utils.PrintAnswer(19, numOfFirstMonthSundays)
}

func GetNumOfDaysSinceStart(month int, year int) int {
	numOfDays := 0

	for i := 1900; i <= year; i++ {
		for j := 1; (i == year && j < month) || (i != year && j <= 12); j++ {
			numOfDays += GetNumOfDaysForMonth(j, i)
		}
	}

	return numOfDays
}

func GetFirstMonthDay(month int, year int) int {

	numOfDaysSinceStart := GetNumOfDaysSinceStart(month, year)
	firstMonthDay := (numOfDaysSinceStart + 1) % 7

	return firstMonthDay
}

func GetNumOfDaysForMonth(month int, year int) int {

	isLeapYear := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	switch month {
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear {
			return 29
		}
		return 28
	}
	return 31
}
