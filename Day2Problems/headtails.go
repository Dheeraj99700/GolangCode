package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}
func main() {
	var num int
	num = rand.Intn(2)
	fmt.Println(num)
	if num%2 == 0 {
		fmt.Printf("Heads\n")
	} else {
		fmt.Printf("Tails\n")
	}
}
