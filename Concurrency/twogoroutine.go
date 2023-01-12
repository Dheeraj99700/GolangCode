package main

import (
	"fmt"
	"sync"
)

var x int

func add(wg *sync.WaitGroup) {
	x := 5
	y := 7
	fmt.Println("Add", x+y)
	wg.Done()

}
func sub(wg *sync.WaitGroup) {
	x := 7
	y := 5
	fmt.Println("Sub", x-y)
	wg.Done()

}

func main() {
	x = 1
	var wg sync.WaitGroup
	wg.Add(2)
	go add(&wg)
	go sub(&wg)
	wg.Wait()

}
