package main 

import "fmt"

func Append() {
	var names []string;

	//append para inserir elementos no final da array
	var addName = append(names, "João")
	addName = append(addName, "Maria")
	addName = append(addName, "Ana", "Mateus", "Rogerio")

	//  fmt.Println(names)
	fmt.Println(addName)
	fmt.Printf("comprimento = %d, capacidade = %d\n", len(addName), cap(addName))
}