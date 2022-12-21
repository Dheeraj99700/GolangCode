package main

import "fmt"

func main() {
	var degree float64 = 0
	var Fahrenheit float64 = 0
	var option int
	fmt.Printf("Select the operation you want to perform \n1.Convert Degree to Fahrenheit\n2.Fahrenheit to degree celcius\n")
	fmt.Scanf("%d", &option)
	switch option {
	case 1:
		{
			fmt.Printf("Enter Temperature in Degree : ")
			fmt.Scanf("%f", &degree)
			Fahrenheit = (degree * 1.8) + 32
			fmt.Printf("Temperature in Fahrenheit is : %0.2f\n", Fahrenheit)
		}
	case 2:
		{
			fmt.Printf("Enter Temperature in Fahrenheit : ")
			fmt.Scanf("%f", &Fahrenheit)
			degree = (Fahrenheit - 32) * 0.5556
			fmt.Printf("Temperature in Fahrenheit is : %0.2f \n", degree)
		}
	default:
		{
			fmt.Println("Invalid input")
		}

	}
}
