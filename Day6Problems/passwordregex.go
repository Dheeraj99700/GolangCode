package main

//Minimum of one small case letter
//Minimum of one upper case letter
//Minimum of one digit
//Minimum of one special character
//Minimum 8 characters length

import (
	"fmt"
	"regexp"
)

func main() {
	var password string
	fmt.Println("Enter the Password")
	fmt.Scanf("%s", &password)
	re := regexp.MustCompile("([!@#$%^&*.?-])+").MatchString(password)

	if re == true {
		fmt.Println("Password is valid")
	} else {
		fmt.Println("Password is INvalid")
	}

}
