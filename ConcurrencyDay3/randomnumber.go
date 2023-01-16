package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(numberchnl chan<- int) {
	for i := 0; i < 15; i++ {
		x := rand.Intn(100)
		numberchnl <- x
		time.Sleep(time.Second)
	}
	close(numberchnl)
}
func main() {
	numberchnl := make(chan int, 3)
	go random(numberchnl)
	for v := range numberchnl {
		fmt.Println(v)
	}
}
