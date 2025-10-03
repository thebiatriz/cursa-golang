package main

import "fmt" 

func main() {
	//float 64 é como o double das outras linguagens
	var height float64 = 1.34
	var open bool = true

	fmt.Printf("Tipo: %T. Valor: %v\n", height, height)
	fmt.Printf("Tipo: %T. Valor: %v\n", open, open)
}