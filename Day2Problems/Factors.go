package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	var factors []string
	fmt.Println("Enter the number")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			str := strconv.Itoa(i)
			factors = append(factors, str)
		}

	}
	fmt.Println(factors)
	for i := 0; i < len(factors); i++ {
		if factors[i] == "3" {
			fmt.Println("Pling")
		} else if factors[i] == "5" {
			fmt.Println("Plang")
		} else if factors[i] == "7" {
			fmt.Println("Plong")
		}

	}
}
