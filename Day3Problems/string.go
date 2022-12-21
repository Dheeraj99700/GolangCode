package main

import (
	"fmt"
	"strings"
)

func main() {

	var p string
	var vowel int
	var cons int
	fmt.Print("Enter the String\n")
	fmt.Scanf("%s", &p)
	_ = strings.Split(p, "")
	for i := 0; i <= len(p)-1; i++ {
		if p[i] == 'a' || p[i] == 'e' || p[i] == 'i' || p[i] == 'o' || p[i] == 'u' ||
			p[i] == 'A' || p[i] == 'E' || p[i] == 'I' || p[i] == 'O' || p[i] == 'U' {
			vowel += 1
		} else {
			cons += 1
		}
	}
	fmt.Printf("Vowels Present in the String are : %d\n", vowel)
	fmt.Printf("Consonant Present in the String are : %d\n", cons)
}
