package main

import "fmt"

func fibonacci(num int) {
	x := 0
	y := 1
	for i := 0; i < num; i++ {
		fmt.Printf("%d\t", x)
		temp := x
		x = y
		y = temp + x
	}

}
func main() {
	var num int
	fmt.Printf("Enter the number till you want fibonacii series\n")
	fmt.Scanf("%d", &num)
	fmt.Printf("The Fibonacci series of number %d is \n", num)
	fibonacci(num)
	fmt.Printf("\n")

}
