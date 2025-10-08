package main

import "fmt"

func FuncInVariable() {

	addFunc := func(number1, number2 float64) float64 {
		return number1 + number2
	}

	multiplyFunc := func(multiplicand, multiplier float64) float64 {
		return multiplicand * multiplier
	}

	fmt.Println("Resultado da soma:", addFunc(3, 4))
	fmt.Println("Resultado da multiplicação:", multiplyFunc(3, 4))

}
