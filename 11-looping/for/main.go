package main

import "fmt"

func main() {
	// somar numeros inteiros de 1 a 10

	var add int = 0

	for i := 1; i <= 10; i++ {
		add += i
	} 

	fmt.Println("Resultado da soma de 1 a 10:", add)

	fmt.Println("\n--------Chamando função de for com range--------")
	Range()
}
