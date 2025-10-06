package main

import "fmt"

func main() {
	// operadores && || ! 

	//	3 < x < 7
	// x > 3 && x < 7

	var x int = 4 
	fmt.Println("x > 3 && x < 5 =", x > 3 && x < 7)

	fmt.Println("x < 3 || x > 7 =", x < 3 || x > 7)

	//invertendo lógica
	fmt.Println(!true)
}
