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
	min := 10
	max := 50
	var num int
	var average int
	var arr = make([]int, 5)
	fmt.Println("The Random Numbers Are")
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(max-min) + min
		num = num + arr[i]
	}
	average = num / 5
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\t", arr[i])
	}
	fmt.Printf("\n")
	fmt.Println("The average of random number is : ", average)

}
