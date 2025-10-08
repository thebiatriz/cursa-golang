package main

import "fmt"

func main() {
	fmt.Println("\n-------Executando a função de armazenar funções em variáveis-------")
	FuncInVariable()

	fmt.Println("\n-------Executando a função de passar funções como argumento-------")
	FuncInVariable()

	fmt.Println("\n-------Executando a função de retorno de funções-------")
	FuncOfFunctions()

	fmt.Println("\n-------Executando a função com closure-------")
	FuncClosure()
}
