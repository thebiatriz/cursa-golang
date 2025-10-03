package main

import "fmt"

func main() {
	var value1 int
	var value2 = value1

	fmt.Printf("Tipo: %T. Valor: %v\n", value2, value2)

	//inferencia com numero complexo
	//parte real + parte imaginaria
	var z = 0.64 + 0.3i
	fmt.Printf("Tipo: %T. Valor: %v\n", z, z)

	//extração da parte real
	realPart := real(z)
	imaginaryPart := imag(z)

	//operações
	w := 2 + 1i
	complexResult := z + w

	fmt.Printf("Parte real do número complexo: %v\n", realPart)
	fmt.Printf("Parte imaginária do número complexo: %v\n", imaginaryPart)
	fmt.Printf("Soma de %v e %v: %v\n", z, w, complexResult)

}
