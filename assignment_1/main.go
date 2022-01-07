package main

import "fmt"

func main() {
	// create a slice of type int which has numbers from 0-10
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// loop through each num in nums
	for _, num := range nums {
		// divide by 2 and check the remainder.
		// If its 0 it's even
		// else its odd
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
