package main

import (
	"fmt"
	"strings"
)

func main() {
	var words string

	fmt.Printf("Enter the string you want to input \n")
	fmt.Scanf("%s", &words)
	x := strings.ContainsAny(words, "a")
	y := strings.ContainsAny(words, "e")
	z := strings.ContainsAny(words, "p")

	if x == true && y == true && z == true {
		fmt.Printf("All Alphabets Present\n")
	} else if x == true || y == true || z == true {
		fmt.Printf("One or more  alphabet Present \n")
	} else {
		fmt.Printf("No alphabet Present\n")
	}

}
