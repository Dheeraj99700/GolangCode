package main

import (
	"fmt"
	"time"
)

func main() {
	year, month, day := time.Now().Date()
	fmt.Println("Day    :", day)
	fmt.Println("Month  :", month)
	fmt.Println("year   :", year)
	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05"))

}
