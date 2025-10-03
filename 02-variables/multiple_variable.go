package main

import "fmt"

func MultipleVariable() {

	//declarando junto com tipos diferentes
	var (
		age    int    = 35
		name   string = "Magalhães"
	)

	// metodo simples com :=
	height := 160

	fmt.Printf("Nome: %s\nIdade: %d\nAltura: %d\n", name, age, height)
}
