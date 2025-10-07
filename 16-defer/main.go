package main

import "fmt"

func print() string {
	fmt.Println("Imprimindo...")
	return "Valor de Imprimir final da função"
}

func main() {
	// defer atrasa a execução até o final de main (da função)
	// depois disso que executa o que está no defer

	//o argumento é executado normalmente, mas o println espera o defer
	defer fmt.Println(print())
	
	defer fmt.Println("Oi")
	defer fmt.Println("Tudo bem?")

	// cada defer vai sendo adicionado como pilha (em cima)

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("Contagem:")

}
