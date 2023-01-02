package main

import (
	"fmt"
)

func filter(numbers []float64, functionName func(float64) float64) []float64 {
	filteredNumbers := make([]float64, len(numbers))
	for i, number := range numbers {
		filteredNumbers[i] = functionName(number)

	}
	return filteredNumbers
}

func Percentage(number float64) float64 {
	return number * 100
}

func halfnumber(number float64) float64 {
	return number / 2
}

func main() {

	numbers := []float64{0.6, 0.3, 0.84, 0.04}

	numbersDecimalpoint := filter(numbers, Percentage)
	fmt.Println(numbersDecimalpoint)

	Halfvalue := filter(numbers, halfnumber)
	fmt.Println(Halfvalue)
}
