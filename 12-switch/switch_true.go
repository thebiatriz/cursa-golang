package main

import "fmt"

func VerifyGrades() {
	var grade int = 7

	switch true {
	case grade > 9: //compara o resultado com a condição do switch
		fmt.Printf("Nota: %d. Resultado: Ótimo", grade)
	case grade > 7:
		fmt.Printf("Nota: %d. Resultado: Muito bem", grade)
	case grade > 6:
		fmt.Printf("Nota: %d. Resultado: Bom", grade)
	default:
		fmt.Printf("Nota: %d. Resultado: Ruim", grade)
	}
}
