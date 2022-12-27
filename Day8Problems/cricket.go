package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var slice []string
	fmt.Println("All Rows DATA")
	data, _ := ioutil.ReadFile("Cricket.csv")
	fmt.Println(string(data))
	readFile, err := os.Open("Cricket.csv")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		slice = append(slice, fileScanner.Text())
	}

	readFile.Close()
	fmt.Println("Data start fron 2nd row")
	for index, values := range slice {
		if index != 0 {
			fmt.Println(values)
		}
	}

}
