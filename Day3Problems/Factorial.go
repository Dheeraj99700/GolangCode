package main

import "fmt"

func factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}

func main() {
	var num int
	fmt.Printf("Enter the number of which you want to find factorial\n")
	fmt.Scanf("%d", &num)
	fmt.Printf("Factorial of number %d is : %d\n", num, factorial(num))

}
