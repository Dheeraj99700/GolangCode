package main

import "fmt"

func odd(slice []int) {
	var odd []int
	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 1 {
			odd = append(odd, slice[i])
		}
	}
	fmt.Printf("The odd numbers are : %d\n", odd)
}
func even(slice []int) {
	var even []int
	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			even = append(even, slice[i])

		}

	}
	fmt.Printf("The even numbers are : %d\n", even)
}

func main() {

	var arr = [20]int{2, 5, 4, 8, 9, 7, 6, 3, 4, 5, 2, 7, 8, 6, 5, 4, 4, 25, 41, 89}
	slice := arr[5:15]
	fmt.Printf("The array is : %d\n", arr)
	fmt.Printf("The Slice is : %d\n", slice)
	odd(slice)
	even(slice)
}
