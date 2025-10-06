package main

import "fmt"

func main() {
	var nome string = "João"

	//tambem pode declarar em switch 
	// switch nome := "João"; nome

	switch nome {
	case "Ana":
		fmt.Println("É a Ana")
	case "João":
		fmt.Println("É o João")
	default:
		fmt.Println("Não conheço")
	}

	fmt.Println("-------Executando a função de verificar as notas-------")
	VerifyGrades()
}
