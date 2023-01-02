package main

import "fmt"

func Average(numbers ...int) float64 {
	add := 0
	for _, val := range numbers {
		add = add + val

	}
	return (float64(add)) / float64((len(numbers)))
}
func main() {
	fmt.Println(Average(1, 2, 3, 4, 5, 6))
}
