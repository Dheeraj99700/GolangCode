package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	var slice []string
	var valid []string
	data, _ := ioutil.ReadFile("LearnerData.txt")
	fmt.Println(string(data))
	readFile, err := os.Open("LearnerData.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		slice = append(slice, fileScanner.Text())
	}

	readFile.Close()
	for index, values := range slice {
		words := strings.Fields(values)
		valid = append(valid, words[4])
		re := regexp.MustCompile("^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$").MatchString(valid[index])
		if re == true {
			x := "Email is Valid"
			words = append(words, x)
		} else {
			y := "Email is Invalid"
			words = append(words, y)
		}
		fmt.Println(words)
	}

}
