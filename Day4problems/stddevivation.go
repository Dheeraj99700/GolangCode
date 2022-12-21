package main

import (
	"fmt"
	"math"
	"strconv"
)

func average(floatnum []float64) float64 {
	var sum float64
	var avg float64
	for i := 0; i < len(floatnum); i++ {
		sum += floatnum[i]
	}
	avg = sum / (float64)(len(floatnum))
	return avg

}

func stddev(floatnum []float64) float64 {
	var sum, avg, deviation float64
	var std float64
	for i := 0; i < len(floatnum); i++ {
		sum += floatnum[i]
	}
	avg = sum / (float64)(len(floatnum))
	for i := 0; i < len(floatnum); i++ {
		std += math.Pow(floatnum[i]-avg, 2)
	}
	deviation = math.Sqrt(std / 10)
	return deviation
}

func main() {
	slice := make([]float64, 3)
	var input string
	var floatnum []float64
	for i := 0; i < len(slice); i++ {
		fmt.Println("Please enter a number")
		fmt.Scanln(&input)
		if input == "Q" {
			break
		}
		if input == "q" {
			break
		}

		myslice, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		slice = append(slice, myslice)

	}
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] != 0 {
			floatnum = append(floatnum, slice[i])
		}

	}

	fmt.Printf("Average of the Float array is : %0.2f\n", average(floatnum))
	fmt.Printf("Deviation of the Float array is : %0.2f\n", stddev(floatnum))

}
