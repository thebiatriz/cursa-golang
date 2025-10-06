package main

import "fmt"

func main() {
	// >; <; >=; <=
	// ==; !=

	var a int = 23
	var b int = 7
	var c float64 = 7

	fmt.Println("a > b:", a > b)
	fmt.Println("b > a:", a < b)
	fmt.Println("b >= a:", b >= a)
	fmt.Println("b <= a:", b <= a)

	fmt.Println("3 == 4:", 3 == 4)
	//fmt.Println("3 == 3:", 3 == 3)

	//não aceita comparações de tipos diferentes, precisa converter
	fmt.Println("b == c:", b == int(c))

	fmt.Println("a != b:", a != b)
	fmt.Println("b != c:", b != int(c))

}
