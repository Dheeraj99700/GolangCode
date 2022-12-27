package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Please Enter the string with comma ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentence := scanner.Text()

	result := strings.ReplaceAll(sentence, ",", ", ")
	fmt.Println(result)

}
