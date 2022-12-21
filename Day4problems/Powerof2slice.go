package main

import "fmt"

func powerof2(inpslice [10]uint8) {
	var poweroftwo []uint8
	for i := 0; i < len(inpslice); i++ {
		if inpslice[i]%2 == 0 {
			poweroftwo = append(poweroftwo, inpslice[i])
		}
	}
	fmt.Printf("The power of 2 slice is : %d\n", poweroftwo)

}

func main() {
	var inpslice [10]uint8
	fmt.Println("Enter the integers to Perform operations")
	for i := 0; i < len(inpslice); i++ {
		fmt.Scanf("%d", &inpslice[i])
	}
	powerof2(inpslice)
}
