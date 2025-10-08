package main

import "fmt"

func compute(fn func(float64, float64) float64) float64 {
	return fn(5, 6)
}

func FuncAsArgument() {

	addFunc := func(number1, number2 float64) float64 {
		return number1 + number2
	}

	multiplyFunc := func(multiplicand, multiplier float64) float64 {
		return multiplicand * multiplier
	}

	fmt.Println("Resultado da soma com argumento:", compute(addFunc))
	fmt.Println("Resultado da multiplicação com argumento:", compute(multiplyFunc))
}
