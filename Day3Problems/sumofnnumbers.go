package main

import "fmt"

func sum(num int) int {

	if num <= 0 {
		return num
	} else {
		if num%3 == 0 || num%5 == 0 {
			return num + sum(num-1)
		} else {
			return sum(num - 1)
		}
	}
}

func main() {
	var num int
	fmt.Printf("Enter the number to calculate the sum of natural numbers before that\n")
	fmt.Scanf("%d", &num)
	fmt.Println(sum(num))

}
