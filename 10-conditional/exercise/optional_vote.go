package main

import (
	"fmt"
)

func main() {
	var age uint

	fmt.Println("Bem-vindo a pesquisa de voto facultativo.")
	fmt.Println("\nInsira a sua idade para visualizar se o seu voto nas eleições nacionais são opcionais ou não:")

	fmt.Scanln(&age)

	if age <= 0 || age > 122 {
		fmt.Println("\nIdade inválida")
	} else if (age >= 16 && age < 18) || age >= 70 {
		fmt.Printf("\nO voto com %d anos é facultativo", age)
	} else if age > 0 && age < 16 {
		fmt.Printf("\nNão é permitido votar com %d ano(s)", age)
	} else {
		fmt.Printf("\nO voto com %d anos é obrigatório", age)
	}
}
