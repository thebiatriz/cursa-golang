package main

import "fmt"

func SimpleArray() {
	// Arrays
	// 1, 1, 2, 3, 5, 8, 13

	var numbers [7]int
	numbers[0] = 1
	numbers[1] = 1
	numbers[2] = 2
	numbers[3] = 3
	numbers[4] = 5
	numbers[5] = 8
	numbers[6] = 13

	fmt.Println(numbers)
}
