package main

// ^  := Means beginning of input
// .  := means "any character"
// ^  :=indicates the beginning of the string
// $  :=indicates the end of the string
import (
	"fmt"
	"regexp"
)

func main() {
	var email string
	fmt.Println("Enter the Email ID")
	fmt.Scanf("%s", &email)
	//str1 := "^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$"
	re := regexp.MustCompile("^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$").MatchString(email)

	if re == true {
		fmt.Println("Email is valid")
	} else {
		fmt.Println("Email is INvalid")
	}

}
