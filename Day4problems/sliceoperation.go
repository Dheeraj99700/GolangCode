package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 3)
	var input string
	var numbers []int
	for i := 0; i < len(slice); i++ {
		fmt.Println("Please enter a number")
		fmt.Scanln(&input)
		if input == "X" {
			break
		}

		myslicevar, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		slice = append(slice, myslicevar)
		if slice[i] != 0 {
			numbers = append(numbers, slice[i])
		}

		sort.Ints(slice)
		fmt.Println(slice)
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] != 0 {
			numbers = append(numbers, slice[i])
		}

	}
	fmt.Println(numbers)

}
