package main

import "fmt"

func main() {
	fmt.Println("\n-------Executando a função para recuperar um slice da array-------")
	TestSlice()

	fmt.Println("\n-------Executando a função de modificar um slice da array-------")
	ModifySlice()

	fmt.Println("\n-------Executando a função de slice literal-------")
	LiteralSlice()

	fmt.Println("\n-------Executando a função de slice com make-------")
	UsingMake()

	fmt.Println("\n-------Executando a função de nil slice-------")
	NilSlice()

	fmt.Println("\n-------Executando a função de append (adicionar elemento)-------")
	Append()
}
