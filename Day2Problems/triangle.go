package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Printf("Enter the length of three side of triangles\n")
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)
	fmt.Scanf("%d", &z)
	if x == y && y == z {
		fmt.Printf("The triangle is equilateral triangle\n")
	} else if x == y || x == z || y == z {

		fmt.Printf("The Triangle is iscosceles triangle\n")
	} else {
		fmt.Printf("The Triangle is scalene triangle\n")
	}

}
