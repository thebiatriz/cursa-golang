package main

import "fmt"

func main() {

	var numbers = [7]int{1, 1, 2, 3, 5, 8, 13}

	var total int = 0

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}

	fmt.Println("\nResultado da soma dos números da array é:", total)
}
