package main

import "fmt"

func SimpleVariable() {

	//declarando o tipo
	var olderAge int = 35

	//inferido. cria e atribui
	currentAge := 20

	var name = "Beatriz"

	fmt.Printf("O meu nome é %s. Eu tenho %d anos e terei %d anos daqui a 15 anos.\n", name, currentAge, olderAge)

}