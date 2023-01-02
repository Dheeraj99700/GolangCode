package main

import (
	"fmt"
	"math"
)

func filter(numbers []float64, functionName func(float64) bool) []float64 {
	filteredNumbers := []float64{}
	for _, number := range numbers {
		if functionName(number) {
			filteredNumbers = append(filteredNumbers, number)
		}
	}
	return filteredNumbers
}

func DecimalPoint(number float64) bool {
	return math.Mod(number, 1) == 0
}

func Divisble(number float64) bool {
	return math.Mod(number, 3) == 0
}

func main() {

	numbers := []float64{1.00, 2.81, 3.00, 5.79, 6.00, 7.41, 8.0, 9.34, 1.79, 11.00, 12.00, 15.00, 5.21, 9.00}

	numbersDecimalpoint := filter(numbers, DecimalPoint)
	fmt.Println(numbersDecimalpoint)

	Divisbleby3 := filter(numbers, Divisble)
	fmt.Println(Divisbleby3)
}
