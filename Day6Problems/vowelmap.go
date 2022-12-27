package main

import (
	"fmt"
	"strings"
)

var VowelsOrConsonats = make(map[string][]string)
var vowela []string
var vowele []string
var voweli []string
var vowelo []string
var vowelu []string
var vowelA []string
var vowelE []string
var vowelI []string
var vowelO []string
var vowelU []string

func vowelOrconsonat(arr []string) (int, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string) {
	var count int = 0
	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] == "a" || arr[i] == "e" || arr[i] == "i" || arr[i] == "o" || arr[i] == "u" ||
			arr[i] == "A" || arr[i] == "E" || arr[i] == "I" || arr[i] == "O" || arr[i] == "U" {
			count += 1
		}
		if arr[i] == "a" {
			vowela = append(vowela, arr[i])
		}
		if arr[i] == "e" {
			vowele = append(vowele, arr[i])
		}
		if arr[i] == "i" {
			voweli = append(voweli, arr[i])
		}
		if arr[i] == "o" {
			vowelo = append(vowelo, arr[i])
		}
		if arr[i] == "u" {
			vowelu = append(vowelu, arr[i])
		}
		if arr[i] == "A" {
			vowelA = append(vowelA, arr[i])
		}
		if arr[i] == "E" {
			vowelE = append(vowelE, arr[i])
		}
		if arr[i] == "I" {
			vowelI = append(vowelI, arr[i])
		}
		if arr[i] == "O" {
			vowelO = append(vowelO, arr[i])
		}
		if arr[i] == "U" {
			vowelU = append(vowelU, arr[i])
		}
	}
	return count, vowela, vowele, voweli, vowelo, vowelu, vowelA, vowelE, vowelI, vowelO, vowelU
}

func main() {

	var p string
	var slice []string

	fmt.Print("Enter the String\n")
	fmt.Scanf("%s", &p)
	arr := strings.Split(p, "")
	for i := 0; i <= len(arr)-1; i++ {
		slice = append(slice, arr[i])
	}
	count, vowela, vowele, voweli, vowelo, vowelu, vowelA, vowelE, vowelI, vowelO, vowelU := vowelOrconsonat(slice)
	VowelsOrConsonats["vowel['a']"] = vowela
	VowelsOrConsonats["vowel['e']"] = vowele
	VowelsOrConsonats["vowel['i']"] = voweli
	VowelsOrConsonats["vowel['o']"] = vowelo
	VowelsOrConsonats["vowel['u']"] = vowelu
	VowelsOrConsonats["vowel['A']"] = vowelA
	VowelsOrConsonats["vowel['E']"] = vowelE
	VowelsOrConsonats["vowel['I']"] = vowelI
	VowelsOrConsonats["vowel['O']"] = vowelO
	VowelsOrConsonats["vowel['U']"] = vowelU

	fmt.Printf("\nCount of Vowel is %d\n", count)
	for keys, values := range VowelsOrConsonats {
		if len(values) != 0 {
			fmt.Printf("%v ---> %v \n", keys, values)
		}
	}

}
