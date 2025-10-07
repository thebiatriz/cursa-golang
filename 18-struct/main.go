package main

import "fmt"

func main() {
	fmt.Println("\n-------Executando a função de definição de struct-------")
	DefineStruct()

	fmt.Println("\n-------Executando a função struct literal-------")
	LiteralStruct()

	fmt.Println("\n-------Executando a função com pointeiro em struct-------")
	PointerStruct()
}
