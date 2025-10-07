package main

import "fmt"

func PointerStruct() {
	type Position struct {
		X int
		Y int
	}

	var pos Position = Position{47, 81}
	var pointer *Position = &pos

	// desreferenciar automaticamente
	pointer.X = 20

	(*pointer).Y = 33

	fmt.Println("Acessando a struct:", pos)
}
