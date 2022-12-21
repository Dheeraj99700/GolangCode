package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(num []int) (int, int) {
	var max int = num[0]
	var secmax int
	for i := 0; i < len(num); i++ {

		if num[i] > max {
			secmax = max
			max = num[i]
		} else if num[i] > secmax {
			secmax = num[i]
		}
	}
	return max, secmax
}
func min(num []int) (int, int) {
	var min int = num[0]
	var secmin int
	for i := 0; i < len(num); i++ {

		if num[i] < min {
			secmin = min
			min = num[i]
		} else if num[i] < secmin {
			secmin = num[i]
		}
	}
	return min, secmin
}

func main() {

	var num []string
	var values []int
	fmt.Println("Please enter comma separated coordinates: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers := scanner.Text()
	odd := strings.Split(numbers, ",") //spliting the string and removing comma
	for i := 0; i < len(odd); i++ {

		num = append(num, odd[i])    //appending the string in string array
		v, _ := strconv.Atoi(num[i]) //converting that value in int
		values = append(values, v)   //storing in int array
	}
	fmt.Printf("\n%d\n", values)
	max(values)
	firstmax, secondmax := max(values)
	fmt.Printf("First Max is  : %d\nSecond Max is : %d\n\n\n", firstmax, secondmax)

	min(values)
	fmt.Println(values)
	firstmin, secondmin := min(values)
	fmt.Printf("First Min is  : %d\nSecond Min is : %d\n", firstmin, secondmin)

}
