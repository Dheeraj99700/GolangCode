package main

import (
	"fmt"
	"strings"
)

func filter[T any](strings []T, functionName func(T) bool) []T {
	filteredstrings := []T{}
	for _, stringvalue := range strings {
		if functionName(stringvalue) {
			filteredstrings = append(filteredstrings, stringvalue)
		}
	}
	return filteredstrings
}

func StringwithB(stringvalue string) bool {
	_ = strings.Split(stringvalue, "")
	if stringvalue[0] == 'b' || stringvalue[0] == 'B' {
		return true
	}
	return false

}

func Stringwith3Char(stringvalue string) bool {
	return len(stringvalue) == 3

}

func main() {

	strings := []string{"ant", "beetle", "Bee", "wasp", "butterfly", "fly", "moth"}

	stringwithB := filter(strings, StringwithB)
	fmt.Println(stringwithB)

	stringwith3Char := filter(strings, Stringwith3Char)
	fmt.Println(stringwith3Char)
}
