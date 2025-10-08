package main

import "fmt"

func main() {
	fmt.Println("\n-------Executando a função de criação de map com make-------")
	CreateMap()

	fmt.Println("\n-------Executando a função de criação de um map literal-------")
	LiteralMap()

	fmt.Println("\n-------Executando a função de map com struct-------")
	MapStruct()
}