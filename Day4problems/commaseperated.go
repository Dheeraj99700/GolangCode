package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var num []string
	fmt.Println("Please enter comma separated coordinates: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers := scanner.Text()
	odd := strings.Split(numbers, ",")
	for i := 0; i < len(odd); i++ {
		num = append(num, odd[i])
	}
	fmt.Println(num)

}
