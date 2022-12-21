package main

import "fmt"

func sum(x, y int) int {

	return (x + y)
}

func diff(x, y int) int {
	return (x - y)

}

func sum_diff(x, y int) (int, int) {
	add := x + y
	sub := x - y
	return add, sub

}

func main() {
	var x, y int
	fmt.Printf("Enter X ")
	fmt.Scanf("%d", &x)
	fmt.Printf("Enter Y ")
	fmt.Scanf("%d", &y)
	fmt.Printf("Addition of two numbers are   : %d\n", sum(x, y))
	fmt.Printf("Difference of two numbers are : %d\n", diff(x, y))
	add, diff := sum_diff(x, y)
	fmt.Printf("Addition : %d\nDifference : %d\n", add, diff)

}
