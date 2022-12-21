package main

import (
	"fmt"
	"sort"
)

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

	keys := make([]string, 0, len(statescapitals))
	for k := range statescapitals {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s --------> %s\n", k, statescapitals[k])
	}
}
