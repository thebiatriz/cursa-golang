package main

import "fmt"

func LiteralStruct() {
	type Position struct {
		X int
		Y int
	}

	var pos Position = Position{Y: 47}
	fmt.Println("Valor de X (indefinido):", pos.X)
	fmt.Println("Valor de Y (definido):", pos.Y)
}
