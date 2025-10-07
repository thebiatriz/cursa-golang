package main

import "fmt"

func UsingMake() {
	// make(tipo, comprimento, capacidade)
	numbers := make([]int, 0, 6)

	fmt.Println(numbers[0:6])
}