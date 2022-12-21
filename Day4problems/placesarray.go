package main

import (
	"fmt"
)

func main() {

	var places [7]string
	var count [7]int
	fmt.Println("Enter your fav places")
	for i := 0; i < len(places); i++ {
		fmt.Scanf("%s", &places[i])
		count[i] = len(places[i])
	}
	for i := 0; i < len(places); i++ {
		fmt.Printf("City name is %s & Character count is : %d\n", places[i], count[i])
	}

}
