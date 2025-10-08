package main

import "fmt"

// map é chave valor

func CreateMap() {
	//a propriedade vai ter um valor string, associado a um int
	var grades map[string]int = make(map[string]int)

	// a chave é o nome
	grades["Ana"] = 9

	grades["Maria"] = 10

	value, exists := grades["Maria"]

	if exists {
		fmt.Println("Nota:", value)
	} else {
		fmt.Println("O valor digitado não existe armazenado")
	}
}
