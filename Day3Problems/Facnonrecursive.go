package main

import "fmt"

func factorial(num int) {
	var fact int = 1
	for i := 1; i <= num; i++ {
		fact *= i
	}
	fmt.Println(fact)

}

func main() {
	var num int

	fmt.Printf("Enter the number of which you want to find factorial\n")
	fmt.Scanf("%d", &num)
	factorial(num)

}
