package main

import "fmt"

func main() {
	var a int = 32

	// pegando o endereço de "a" e armazenando no ponteiro
	var p *int = &a

	fmt.Println("Endereço de memória de a armazenado em p:", p)
	fmt.Println("Valor de p:", *p)

	// desreferenciando e atribuindo para 18
	*p = 18

	fmt.Println("Novo valor de p depois de desreferenciar e atribuir:", *p)
}