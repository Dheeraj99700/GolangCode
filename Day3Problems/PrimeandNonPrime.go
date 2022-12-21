package main

import "fmt"

func NonPrimeOdd(num int) {
	count := 0
	for i := num - 1; i > 1; i-- {
		ans := IsNumPrime(i)
		if ans == false {
			if i%2 != 0 {
				fmt.Printf("Non Prime Odd is: %d\n", i)
				count++
			}
		}
	}
	fmt.Printf("Count is :%d ", count)
}
func IsNumPrime(num int) bool {
	is_Prime := 0
	if num == 0 {

	} else {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {

				is_Prime = 1
				return false
			}
		}
		if is_Prime == 0 {
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Println("Enter number: ")
	fmt.Scan(&n)
	NonPrimeOdd(n)
}
