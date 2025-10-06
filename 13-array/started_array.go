package main

import "fmt"

func StartedArray() {
	var numbers1 = [7]int{1, 1, 2, 3, 5, 8, 13}
	numbers2 := []int{1, 1, 2, 3, 5, 8, 13, 21}

	fmt.Println(numbers1)
	fmt.Println(numbers2)
}
