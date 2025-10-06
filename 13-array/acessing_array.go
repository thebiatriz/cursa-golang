package main

import "fmt"

func AcessingArray() {

	numbers := []int{1, 1, 2, 3, 5, 8, 13}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
