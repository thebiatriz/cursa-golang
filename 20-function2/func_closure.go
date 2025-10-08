package main

import "fmt"

func additive() func(int) int {
	var add int = 0

	//closure - referencia variáveis de fora da função e reatribui
	return func(value int) int {
		add += value
		return add
	}
}

func FuncClosure() {
	var ad1 = additive()

	fmt.Println(ad1(21))
	fmt.Println(ad1(21))

	fmt.Println("-----------------")

	var ad2 = additive()
	fmt.Println(ad2(7))
	fmt.Println(ad2(0))
}
