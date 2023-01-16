package main

import (
	"fmt"
)

func main() {
	var size int
	var number int
	evenchnl := make(chan int)
	oddchnl := make(chan int)
	fmt.Println("Enter the number of values you want to input")
	fmt.Scan(&size)
	var input []int
	fmt.Println("Enter the values :")
	for i := 0; i < size; i++ {
		fmt.Scan(&number)
		input = append(input, number)
	}
	go func() {
		for _, value := range input {
			if value%2 == 0 {
				evenchnl <- value
			}
		}
		close(evenchnl)

	}()
	go func() {
		for _, value := range input {
			if value%2 != 0 {
				oddchnl <- value
			}
		}
		close(oddchnl)
	}()
	var evennumbers []int
	var oddnumbers []int
	for value := range evenchnl {
		evennumbers = append(evennumbers, value)
	}
	for value := range oddchnl {
		oddnumbers = append(oddnumbers, value)
	}

	fmt.Println("Even numbers are : ", evennumbers)
	fmt.Println("Odd numbers are : ", oddnumbers)
}
