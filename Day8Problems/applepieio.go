package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonInput := `{
			"ThinlySlicedPeeledApples" : "6 Cups",
			"Sugar" : "3/4 cup",
			"Flour" : "2 tablespoons",
			"Cinamon" : "1/4 teaspoon",
			"Rutmeg" : "1/8 tablespoon",
			"Lemonjuice": "1 tablespoon"
		   }`

	var Applepierecipe = make(map[string]string)
	err := json.Unmarshal([]byte(jsonInput), &Applepierecipe)
	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}
	file, _ := os.Create("ApplePie.text")

	for key, value := range Applepierecipe {
		fmt.Printf("%v --------> %v\n", key, value)
		_, _ = file.WriteString(key + "--------->" + value + "\n")

	}

}
