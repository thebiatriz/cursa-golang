package main

import "fmt"

func Range() {
	numbers := []int{10, 20, 30, 40}

	for index, valor := range numbers {
		fmt.Printf("Indíce %d: %d\n", index, valor)
	}
}
