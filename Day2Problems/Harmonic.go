package main

import (
	"fmt"
)

func main() {
	var harmonic float32 = 1.00
	var num int
	fmt.Printf("Enter till the n value you want to find harmoic number \n")
	fmt.Scanf("%d", &num)
	for i := 2; i <= num; i++ {
		harmonic += 1.00 / (float32)(i)
	}
	fmt.Printf("Harmoic number :%0.5f\n", harmonic)

}
