package main

import (
	"fmt"
	"math"
)

func main() {
	var P, Y, R, emi float64
	fmt.Printf("Enter the Principal Amount of Loan\n")
	fmt.Scanf("%f", &P)
	fmt.Printf("Enter the tenure in years of Loan\n")
	fmt.Scanf("%f", &Y)
	fmt.Printf("Enter the Interest Rate  of Loan\n")
	fmt.Scanf("%f", &R)
	R = R / (12 * 100) //One monthinterest
	Y = Y * 12         //one month period
	emi = (P * R * math.Pow(1+R, Y)) / (math.Pow(1+R, Y) - 1)

	fmt.Printf("Monthly Emi of the student loan is : %0.2f\n", emi)
}
