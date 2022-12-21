package main

import "fmt"

func main() {
	var dimarr = [5][2]int{{10}, {12}, {15}, {19}, {24}}
	var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, dimarr[i][j])
			temp := dimarr[i][j]
			x := temp * 2
			dimarr[i][1] = x

		}
	}

}
