package main

import (
	"fmt"
	"strings"
)

func filtermap(strSlice []string, functioname func(string) string) []string {

	filteredNumbers := make([]string, len(strSlice))

	for i, word := range strSlice {
		filteredNumbers[i] = functioname(word)
	}
	return filteredNumbers
}
func FirstCapital(word string) string {
	return strings.Title(word)
}
func Addcolon(word string) string {
	return word + ":"
}

func main() {
	strSlice := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	Capital := filtermap(strSlice, FirstCapital)
	fmt.Println(Capital)
	Colon := filtermap(strSlice, Addcolon)
	fmt.Println(Colon)

}
