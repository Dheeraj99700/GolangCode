package main

import (
	"fmt"
	"math"
)

func main() {
	var p float64
	var n float64
	var op int
	fmt.Println("Enter the principal amount")
	fmt.Scanf("%f", &p)
	fmt.Println("Enter Number of years")
	fmt.Scanf("%f", &n)
	fmt.Printf("Enter the choice you want to calculate \n1.Low Risk\n2.Medium Risk\n3.High Risk\n")
	fmt.Scanf("%d", &op)
	switch op {
	case 1:
		r := 6.25
		//x := p * math.Pow((1+(r/q)), n*q)
		//x := p * (1 + r/n) * (n * q)
		x := p * (math.Pow((1 + r/100), n))
		fmt.Printf("Final Amount is : %0.2f\n", x)

		break
	case 2:
		r := 7.50
		x := p * (math.Pow((1 + r/100), n))
		fmt.Printf("Final Amount is : %0.2f\n", x)
		break
	case 3:
		r := 9.5
		x := p * (math.Pow((1 + r/100), n))
		fmt.Printf("Final Amount is : %0.2f\n", x)
		break
	default:
		fmt.Println("Invalid choice")
		break
	}

}
