package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var stake int
	var goal int
	var trials int
	var bets, wins int = 0, 0
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Enter stake")
	fmt.Scanf("%d", &stake)
	fmt.Println("Enter goal")
	fmt.Scanf("%d", &goal)
	fmt.Println("Enter trails")
	fmt.Scanf("%d", &trials)
	for t := 0; t < trials; t++ {
		for stake > 0 && stake < goal {
			bets++
			if rand.Float64() < 0.5 {
				stake++
			} else {
				stake-- // lose $1
			}
			if stake == goal {
				wins++
			}
		}
	}
	fmt.Printf("%d wins of %d \n", wins, trials)
	fmt.Printf("Percent of games of %f\n", 100*(float64)(wins)/(float64)(trials))
	fmt.Printf("Average of bets are %f\n", 1.0*(float64)(bets)/(float64)(trials))

}
