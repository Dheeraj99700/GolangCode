package main

import "fmt"

func main() {
	var statescapitals = make(map[string]string)
	statescapitals["Maharashtra"] = "Mumbai"
	statescapitals["Goa"] = "Panaji"
	statescapitals["Gujarat"] = "Gandhinagar"
	statescapitals["Jharkhand"] = "Ranchi"
	statescapitals["Karnataka"] = "Bengaluru"
	statescapitals["Kerala"] = "Thiruvananthapuram"
	statescapitals["Manipur"] = "Imphal"
	statescapitals["Odisha"] = "Bhubaneswarr"
	statescapitals["Punjab"] = "Chandigarh"
	statescapitals["Rajasthan"] = "Jaipur"
	fmt.Printf("%T\n", statescapitals)

	for states, capitals := range statescapitals {
		fmt.Printf("%s -------→ %s \n", states, capitals)
	}
}
