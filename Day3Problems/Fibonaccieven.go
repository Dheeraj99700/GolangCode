package main

import "fmt"

func fib(fibonacci []int) {
	var count = 0
	for i := 0; i < len(fibonacci); i++ {
		if fibonacci[i]%2 == 0 {
			count += 1
		}
	}
	fmt.Printf("The count of even numbers are : %d\n", count)
}
func main() {
	x := 0
	y := 1
	var fibonacci []int
	var num int
	fmt.Printf("Enter the number till you want fibonacii series\n")
	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		fibonacci = append(fibonacci, x)
		temp := x
		x = y
		y = temp + x
	}
	fmt.Printf("The fibonacci series till the number %d is : %d\n", num, fibonacci)
	fib(fibonacci)

}
