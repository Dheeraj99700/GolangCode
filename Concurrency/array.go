package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {

	var input = []float64{2.5, 1.4, 8.5, 9.6, 5.4}
	p1 := input[:len(input)/4]
	p2 := input[len(input)/4 : 2*len(input)/4]
	p3 := input[2*len(input)/4 : 3*len(input)/4]
	p4 := input[3*len(input)/4:]
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		sort.Float64s(p1)
		wg.Done()
	}()
	go func() {
		sort.Float64s(p2)
		wg.Done()
	}()
	go func() {
		sort.Float64s(p3)
		wg.Done()
	}()
	go func() {
		sort.Float64s(p4)
		wg.Done()
	}()
	wg.Wait()
	var outputarray []float64
	outputarray = append(outputarray, p1...)
	outputarray = append(outputarray, p2...)
	outputarray = append(outputarray, p3...)
	outputarray = append(outputarray, p4...)
	fmt.Println(outputarray)
}
