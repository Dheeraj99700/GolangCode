package main

import (
	"fmt"
	"math/rand"
	"time"
)

var num1to5 int
var num5to10 int
var num11to15 int
var num16to20 int
var num21to25 int
var num1to5slice []int
var num5to10slice []int
var num11to15slice []int
var num16to20slice []int
var num21to25slice []int
var numberdistribution = make(map[string][]int)

func main() {
	var numbers1 [1000]int
	var num1to5count int = 0
	var num5to10count int = 0
	var num11to15count int = 0
	var num16to20count int = 0
	var num21to25count int = 0

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(numbers1); i++ {
		numbers1[i] = rand.Intn(25)
		if numbers1[i] > 0 && numbers1[i] <= 5 { //condition to find numbers between 1 to 5
			num1to5count += 1
			num1to5slice = append(num1to5slice, numbers1[i])
		}
		if numbers1[i] > 5 && numbers1[i] <= 10 { //condition to find numbers between 6 to 10
			num5to10count += 1
			num5to10slice = append(num5to10slice, numbers1[i])
		}
		if numbers1[i] > 10 && numbers1[i] <= 15 { //condition to find numbers between 11 to 15
			num11to15count += 1
			num11to15slice = append(num11to15slice, numbers1[i])
		}
		if numbers1[i] > 15 && numbers1[i] <= 20 { //condition to find numbers between 16 to 20
			num16to20count += 1
			num16to20slice = append(num16to20slice, numbers1[i])
		}
		if numbers1[i] > 20 && numbers1[i] <= 25 { //condition to find numbers between 21 to 25
			num21to25count += 1
			num21to25slice = append(num21to25slice, numbers1[i])
		}

	}

	numberdistribution[">0  && <=5 "] = num1to5slice
	numberdistribution[">5  && <=10"] = num5to10slice
	numberdistribution[">10 && <=15"] = num11to15slice
	numberdistribution[">15 && <=20"] = num16to20slice
	numberdistribution[">20 && <=25"] = num21to25slice
	for condition, counts := range numberdistribution {
		fmt.Printf("%v ---> %v \n", condition, counts)
	}

}
