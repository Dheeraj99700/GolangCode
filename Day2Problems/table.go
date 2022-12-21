package main

import "fmt"

func main() {
	var num int
	var arr = make([]int, 11)
	fmt.Printf("Enter the number between 2 to 25 which table you want to print \n")
	fmt.Scanf("%d", &num)
	fmt.Printf("Table of %d is :\n", num)
	if num >= 2 && num <= 25 {
		for i := 1; i <= 10; i++ {
			arr[i] = num * i
		}
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d*%d=%d\n", num, i, arr[i])
		}

	} else {
		fmt.Printf("Invalid Number table not available")
	}

}
