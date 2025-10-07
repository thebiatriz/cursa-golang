package main

import "fmt"

func TestSlice() {
	var numbers = [7]int{1, 1, 2, 3, 5, 8, 13}

	// Slice
	fmt.Println("Slice da array:", numbers[2:5]) //inclui até o 5

	// Omitindo o final = pega a partir do índice 2 até o final
	fmt.Println("Omitindo o final (começa do indíce 2):", numbers[2:])

	// Omitindo o começo = pega do começo até o indíce 5
	fmt.Println("Omitindo o começo (vai até o indíce 5):", numbers[:5])
}
