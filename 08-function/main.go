package main

import "fmt"

func main() {
	squared1, cubed1 := Calc(2)
	squared2, cubed2 := CalcNamedReturn(4)

	fmt.Println("Resultado da soma:", Add(10, 15))
	fmt.Println("Resultado do cálculo (sem retorno nomeado):", squared1, cubed1)
	fmt.Println("Resultado do cálculo (sem retorno nomeado):", squared2, cubed2)
}
