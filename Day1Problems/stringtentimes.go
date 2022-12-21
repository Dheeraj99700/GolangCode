package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	var num1 int

	fmt.Printf("Enter the number in string you want to print 10 times of that number \n")
	fmt.Scanf("%s", &num) //input as string
	re, _ := strconv.Atoi(num)
	//string to int coversion
	num1 = re * 10                  //multiply
	str := strconv.Itoa(num1)       //conerting result of type int into string
	fmt.Printf("%s %T\n", str, str) //printing string output
	//fmt.Printf(str)

}
