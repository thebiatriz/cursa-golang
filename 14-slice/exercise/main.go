package main

import (
	"fmt"
)

func main() {
	var animes = [3]string{"Dragon Ball", "Sailor Moon", "Death Note"}

	var firstTwo = animes[:2]
	var lastTwo = animes[1:]

	fmt.Println(animes)
	fmt.Println("Dois primeiros elementos da Array:", firstTwo)
	fmt.Printf("comprimento = %d, capacidade = %d\n", len(firstTwo), cap(firstTwo))

	fmt.Println("\n------------------------")

	fmt.Println("\nDois últimos elementos da Array:", lastTwo)
	fmt.Printf("comprimento = %d, capacidade = %d\n", len(lastTwo), cap(lastTwo))

}
