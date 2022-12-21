package main

import (
	"fmt"
)

func main() {
	var firstname string
	var lastname string
	fmt.Printf("Enter FirstName  OF a person : ")
	fmt.Scanf("%s", &firstname)
	fmt.Printf("Enter LastName OF a person : ")
	fmt.Scanf("%s", &lastname)
	fmt.Printf("The FullName of person is :%s %s\n", firstname, lastname)

}
