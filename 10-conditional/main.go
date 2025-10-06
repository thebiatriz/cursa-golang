package main

import "fmt"

func main() {
	// if  else  else if

	var a int = 11

	if a > 10 {
		fmt.Println("a é maior que 10. valor de a:", a)
	} else if a > 3 {
		fmt.Println("a está entre 4 e 10. valor de a:", a)
	} else {
		fmt.Println("a é menor ou igual a 5. valor de a:", a)
	}


	fmt.Println("\nExecutando função com variável definida no bloco condicional:")
	variableInBlock()
}

func variableInBlock() {
	if a := 20; a > 20 {
		fmt.Println("a é maior que 20. Valor de a:", a)
	} else if a > 5 {
		fmt.Println("a está entre 6 e 20. Valor de a:", a)
	} else {
		fmt.Println("a é menor ou igual a 5. Valor de a:", a)
	}
}
