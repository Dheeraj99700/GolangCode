package main

import (
	"fmt"
	"sort"
)

func main() {
	var size int
	var number float64
	chnl := make(chan []float64)
	fmt.Println("Enter the number of values you want to input")
	fmt.Scan(&size)
	var input []float64
	fmt.Println("Enter the values :")
	for i := 0; i < size; i++ {
		fmt.Scan(&number)
		input = append(input, number)
	}
	p1 := input[:len(input)/4]
	p2 := input[len(input)/4 : 2*len(input)/4]
	p3 := input[2*len(input)/4 : 3*len(input)/4]
	p4 := input[3*len(input)/4:]
	go func() {
		sort.Float64s(p1)
		fmt.Println("Sorted p1 is : ", p1)
		chnl <- p1

	}()
	go func() {
		sort.Float64s(p2)
		fmt.Println("Sorted p2 is : ", p2)
		chnl <- p2
	}()
	go func() {
		sort.Float64s(p3)
		fmt.Println("Sorted p3 is : ", p3)
		chnl <- p3
	}()
	go func() {
		sort.Float64s(p4)
		fmt.Println("Sorted p4 is : ", p4)
		chnl <- p4
	}()
	output1 := <-chnl
	output2 := <-chnl
	output3 := <-chnl
	output4 := <-chnl
	var outputarray []float64
	outputarray = append(outputarray, output1...)
	outputarray = append(outputarray, output2...)
	outputarray = append(outputarray, output3...)
	outputarray = append(outputarray, output4...)
	sort.Float64s(outputarray)
	fmt.Println(outputarray)
}
