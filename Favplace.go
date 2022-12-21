package main

import (
	"fmt"
	"strings"
)

func main() {
	var Favplace string = "Pune"
	fmt.Printf("My Fav Place to visit is : %s\n", strings.ToUpper(Favplace))
	fmt.Println("My fav place to visit is : ", strings.ToLower(Favplace))
}
